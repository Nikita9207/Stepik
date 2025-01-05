//Количество нулей
//По данным числам, определите количество чисел, которые равны нулю.
//Входные данные
//Вводится натуральное число N, а затем N чисел.
//Выходные данные
//Выведите количество чисел, которые равны нулю.

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

	count := 0

	for _, v := range num {
		if v == 0 {
			count++
		}
	}

	fmt.Println(count)
}
