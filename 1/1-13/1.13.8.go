//Количество минимумов
//Найдите количество минимальных элементов в последовательности.
//Входные данные
//Вводится натуральное число N, а затем N целых чисел последовательности.
//Выходные данные
//Выведите количество минимальных элементов последовательности.

package main

import (
	"fmt"
)

func main() {
	var N, n int
	fmt.Scan(&N)

	num := make([]int, N)

	for i := 0; i < len(num); i++ {
		fmt.Scan(&n)
		num[i] = n
	}

	min := num[0]

	for _, v := range num {
		if v < min {
			min = v
		}
	}

	count := 0

	for _, v := range num {
		if v == min {
			count++
		}
	}

	fmt.Println(count)
}
