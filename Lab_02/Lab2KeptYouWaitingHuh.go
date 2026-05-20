package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct{ Radius float64 }
type Rectangle struct{ Width, Height float64 }
type Triangle struct{ A, B, C float64 }

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

func (t Triangle) Area() float64 {
	p := t.Perimeter() / 2
	return math.Sqrt(p * (p - t.A) * (p - t.B) * (p - t.C))
}
func (t Triangle) Perimeter() float64 { return t.A + t.B + t.C }

// ==========================================
func main() {
	// --- 1 зав
	var a [10]int
	b := []int{5, 12, 7, 3, 9, 1, 14, 8, 2, 6} // Довільні значення для слайсу
	result := make([]int, 10)

	for i := 0; i < 10; i++ {
		a[i] = i + 1 // Значення від 1 до 10
		result[i] = a[i] + b[i]
	}

	fmt.Println("=== Масиви та слайси ===")
	fmt.Println("Масив a:", a)
	fmt.Println("Слайс b:", b)
	fmt.Println("Слайс result (сума):", result)

	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 4, Height: 6},
		Triangle{A: 3, B: 4, C: 5},
	}

	fmt.Println("\n=== Інтерфейси та Структури ===")
	for i, shape := range shapes {
		fmt.Printf("Фігура %d | Площа: %.2f | Периметр: %.2f\n", i+1, shape.Area(), shape.Perimeter())
	}
}
