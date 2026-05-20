package main

import (
	"fmt"
)

func generate() <-chan int {
	out := make(chan int, 100)

	go func() {
		for i := 1; i <= 100; i++ {
			out <- i
		}

		close(out)
	}()

	return out
}

func filterEven(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {

		for n := range in {
			if n%2 == 0 {
				out <- n
			}
		}
		close(out)
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}

func sum(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		total := 0
		for n := range in {
			total += n
		}
		out <- total
		close(out)
	}()

	return out
}

func main() {
	fmt.Println("=== Запуск Конвеєра (Pipeline) ===")

	stage1Chan := generate()
	stage2Chan := filterEven(stage1Chan)
	stage3Chan := square(stage2Chan)
	finalChan := sum(stage3Chan)

	result := <-finalChan

	fmt.Printf("Фінальна сума квадратів усіх парних чисел від 1 до 100: %d\n", result)
	fmt.Println("=== Конвеєр успішно завершив роботу ===")
}
