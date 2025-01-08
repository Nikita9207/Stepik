//Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.
//Входные данные
//Вводится четыре числа.
//Выходные данные
//Необходимо вернуть из функции наименьшее из 4-х данных чисел

package main

import "fmt"

func main() {
	var a int

	a = minimumFromFour()
	fmt.Print(a)
}

func minimumFromFour() int {
	var x int

	var arr = [4]int{}

	for i := 0; i < 4; i++ {
		fmt.Scan(&x)
		arr[i] = x
	}

	x = arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < x {
			x = arr[i]
		}
	}

	return x
}
