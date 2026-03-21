package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("-+- Якщо х у діапазоні (10, 100), то Ф2, інакше Ф1 -+-")
	x := float64(rand.Intn(120))

	var ySwitch float64

	fmt.Printf("Згенероване випадкове значення x: %.2f\n\n", x)

	switch {
	case x > 10 && x < 100:
		ySwitch = (x - 3) / (4 + x)
		fmt.Println("Використано Функцію 2")
	default:
		ySwitch = math.Pow(x, 3) - math.Pow(x, 2) - x
		fmt.Println("Використано Функцію 1")
	}

	fmt.Printf("Результат розрахунку y: %.4f\n", ySwitch)
}
