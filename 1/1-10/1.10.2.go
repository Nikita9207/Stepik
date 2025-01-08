//Требуется написать программу, при выполнении которой с клавиатуры считываются два натуральных числа A и B
//(каждое не более 100, A < B). Вывести сумму всех чисел от A до B  включительно.

package main

import (
	"fmt"
)

func main() {
	var a, b int
	var sum int
	fmt.Scan(&a, &b)

	if a > 100 || b > 100 {
		fmt.Println("Введите число меньше 100")
	}

	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Print(sum)
}
