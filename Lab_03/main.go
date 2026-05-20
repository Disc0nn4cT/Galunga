package main

import (
	"fmt"
	"simplecalc/calc"
)

// ==========================================

type Calculator interface {
	Sum(nums ...float64) float64
	Max(nums ...float64) float64
	Min(nums ...float64) float64
	Divide(a, b float64) (float64, error)
}

type Calc struct{}

func (c Calc) Sum(nums ...float64) float64          { return calc.Sum(nums...) }
func (c Calc) Max(nums ...float64) float64          { return calc.Max(nums...) }
func (c Calc) Min(nums ...float64) float64          { return calc.Min(nums...) }
func (c Calc) Divide(a, b float64) (float64, error) { return calc.Divide(a, b) }

// ==========================================
func main() {
	fmt.Println("\n=== Завдання 1: Використання пакету calc ===")

	fmt.Printf("Сума (1, 2, 3, 4, 5): %.2f\n", calc.Sum(1, 2, 3, 4, 5))
	fmt.Printf("Максимум (10, -5, 42, 7): %.2f\n", calc.Max(10, -5, 42, 7))
	fmt.Printf("Мінімум (10, -5, 42, 7): %.2f\n", calc.Min(10, -5, 42, 7))

	res, err := calc.Divide(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Ділення (10 / 2): %.2f\n", res)
	}

	_, errZero := calc.Divide(5, 0)
	if errZero != nil {
		fmt.Println("Спроба ділення (5 / 0):", errZero)
	}

	fmt.Println("\n=== Завдання 2: Використання інтерфейсу Calculator ===")

	myCalculator := Calc{}

	fmt.Printf("Сума через інтерфейс (10, 20): %.2f\n", myCalculator.Sum(10, 20))
	fmt.Printf("Максимум через інтерфейс (100, 200, 50): %.2f\n", myCalculator.Max(100, 200, 50))
}
