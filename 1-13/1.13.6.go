//Даны два числа. Найти их среднее арифметическое.
//Формат входных данных
//На вход дается два целых положительных числа a и b.
//Формат выходных данных
//Программа должна вывести среднее арифметическое чисел a и b (ответ может быть целым числом или дробным)

package main

import (
	"fmt"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)

	sum := (a + b) / 2
	fmt.Print(sum)
}
