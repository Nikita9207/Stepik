//По данному целому числу, найдите его квадрат.
//Формат входных данных
//На вход дается одно целое число.
//Формат выходных данных
//Программа должна вывести квадрат данного числа.

package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	a = a * a
	fmt.Print(a)
}
