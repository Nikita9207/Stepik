//Вклад в банке составляет x рублей. Ежегодно он увеличивается на p процентов, после чего дробная часть копеек отбрасывается.
//Каждый год сумма вклада становится больше. Определите, через сколько лет вклад составит не менее y рублей.
//Входные данные
//Программа получает на вход три натуральных числа: x, p, y.
//Выходные данные
//Программа должна вывести одно целое число.

package main

import (
	"fmt"
)

func main() {
	var x, p, y, count int

	fmt.Scan(&x)
	fmt.Scan(&p)
	fmt.Scan(&y)

	for x < y {
		x = x + x*p/100
		count++
	}
	fmt.Println(count)
}
