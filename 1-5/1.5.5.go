//Дано неотрицательное целое число. Найдите число десятков (то есть вторую цифру справа).
//Формат входных данных
//На вход дается натуральное число, не превосходящее 10000.
//Формат выходных данных
//Выведите одно целое число - число десятков.

package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	a = (a / 10) % 10
	fmt.Print(a)
}
