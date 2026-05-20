package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	task1Mutex()
	task2Atomic()
}

// ==========================================
func task1Mutex() {
	fmt.Println("=== Завдання 1: М'ютекси (sync.Mutex) ===")

	evenCh := make(chan int)
	oddCh := make(chan int)

	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {

			select {
			case val, ok := <-evenCh:
				if !ok {
					return
				}
				if val%3 == 0 {
					mu.Lock()
					counter++
					mu.Unlock()
				}
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case val, ok := <-oddCh:
				if !ok {
					return
				}
				if val%33 == 0 {
					mu.Lock()
					counter--
					mu.Unlock()
				}
			}
		}
	}()

	for i := 1; i <= 1000; i++ {
		if i%2 == 0 {
			evenCh <- i
		} else {
			oddCh <- i
		}
	}

	close(evenCh)
	close(oddCh)

	wg.Wait()

	fmt.Printf("Фінальне значення counter (Mutex): %d\n", counter)
}

// ==========================================
func task2Atomic() {
	fmt.Println("\n=== Завдання 2: Атомарні операції (sync/atomic) ===")

	evenCh := make(chan int)
	oddCh := make(chan int)

	var counter int64
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case val, ok := <-evenCh:
				if !ok {
					return
				}
				if val%3 == 0 {

					atomic.AddInt64(&counter, 1)
				}
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case val, ok := <-oddCh:
				if !ok {
					return
				}
				if val%33 == 0 {

					atomic.AddInt64(&counter, -1)
				}
			}
		}
	}()

	for i := 1; i <= 1000; i++ {
		if i%2 == 0 {
			evenCh <- i
		} else {
			oddCh <- i
		}
	}

	close(evenCh)
	close(oddCh)
	wg.Wait()

	fmt.Printf("Фінальне значення counter (Atomic): %d\n", atomic.LoadInt64(&counter))
}
