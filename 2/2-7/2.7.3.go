//Дана строка, содержащая только арабские цифры. Найти и вывести наибольшую цифру.
//
//Входные данные
//Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков и строка содержит только арабские цифры.
//
//Выходные данные
//Выведите максимальную цифру, которая встречается во введенной строке.
//
//Sample Input:
//1112221112
//Sample Output:
//2

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n string
	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Println("Ошибка!")
	}

	var nums []int

	for _, v := range n {
		m, _ := strconv.Atoi(string(v))
		nums = append(nums, m)
	}

	x := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] > x {
			x = nums[i]
		}
	}
	fmt.Printf("%d", x)
}
