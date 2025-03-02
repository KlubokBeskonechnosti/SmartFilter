package main

import (
	"fmt"
	"math"
)

const (
	Even     = "even"
	Odd      = "odd"
	Positive = "positive"
	Negative = "negative"
	Squares  = "squares"
	Prime    = "prime"
	Factor10 = "factor10"
)


func filter(slice []int, filterFunc func(int) bool) []int {
	// менять эту функцию не нужно!!!
	// но я советую вам внимательно изучить эту функцию для саморазвития!

	// эта функция проходится по каждому элементу слайса
	// и отфильтровывает те, для которых filterFunc возвращает false

	result := make([]int, 0, len(slice))

	for _, num := range slice {
		if filterFunc(num) {
			result = append(result, num)
		}
	}
	return result
}


func main() {
	// дополните эту мапу своими анонимными функциями!
	filters := map[string]func(int) bool{
		Even: func(n int) bool{
			return n % 2 == 0
		},   // Только чётные числа

		Odd: func(n int) bool {
			return n % 2 != 0
		},   // Только нечётные числа

		Positive: func(n int) bool {
			return n > 0
		},   // Только положительные числа

		Negative: func(n int) bool {
			return n < 0
		},   // Только отрицательные числа

		Squares:  func(n int) bool {
			sqrt := int(math.Sqrt(float64(n)))
			return sqrt*sqrt == n
		},   // Только числа, которые являются квадратами

		Prime: func(n int) bool {
			if n < 2 {
				return false
			}
			for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
				if n%i == 0 {
					return false
				}
			}
			return true
		
		},   // Только простые числа

		Factor10: func(n int) bool {
		return n % 10 == 0 // Только числа, которые кратны 10
	},
}

	// Ниже ничего менять не нужно
	var n int
	fmt.Scan(&n)

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	var operation string
	fmt.Scan(&operation)

	if filterFunc, exists := filters[operation]; exists {
		result := filter(numbers, filterFunc)
		for _, num := range result {
			fmt.Print(num, " ")
		}
	} 
}
