//На вход подается целое число. Необходимо возвести в квадрат каждую цифру числа и вывести получившееся число.
//
//Например, у нас есть число 9119. Первая цифра - 9. 9 в квадрате - 81. Дальше 1. Единица в квадрате - 1. В итоге получаем 811181
//
//Sample Input:
//9119
//Sample Output:
//811181

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

	result := ""

	for _, v := range n {
		m, _ := strconv.Atoi(string(v))
		m *= m
		result += strconv.Itoa(m)
	}

	fmt.Printf("%s", result)
}
