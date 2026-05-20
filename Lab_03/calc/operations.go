package calc

import (
	"errors"
	"fmt"
)

func init() {
	fmt.Println("[Система] Пакет 'calc' успішно ініціалізовано!")
}

func Sum(nums ...float64) float64 {
	sum := 0.0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func Max(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	maxVal := nums[0]
	for _, n := range nums {
		if n > maxVal {
			maxVal = n
		}
	}
	return maxVal
}

func Min(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	minVal := nums[0]
	for _, n := range nums {
		if n < minVal {
			minVal = n
		}
	}
	return minVal
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {

		return 0, errors.New("помилка: ділення на нуль неможливе")
	}
	return a / b, nil
}
