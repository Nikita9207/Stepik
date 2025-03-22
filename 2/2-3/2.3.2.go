//Поменяйте местами значения переменных на которые ссылаются указатели. После этого переменные нужно вывести.
//ВАЖНО: Считайте что пакет main уже объявлен, а также функция main() уже есть.

package main

import "fmt"

func main() {
	var x1, x2 int
	x1 = 2
	x2 = 4
	fmt.Println(x1, x2)
	test(&x1, &x2)
}

func test(x1 *int, x2 *int) {
	*x1, *x2 = *x2, *x1
	fmt.Print(*x1, *x2)
}
