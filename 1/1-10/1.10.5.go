//Найдите первое число от 1 до n включительно, кратное c, но НЕ кратное d.
//Входные данные
//Вводится 3 натуральных числа n, c, d, каждое из которых не превышает 10000.
//Выходные данные
//Вывести первое число от 1 до n включительно, кратное c, но НЕ кратное d. Если такого числа нет - выводить ничего не нужно.

package main

import "fmt"

func findFirstNumber(n, c, d int) int {
	for i := 1; i <= n; i++ {
		if i%c == 0 && i%d != 0 {
			return i
		}
	}
	return 0
}

func main() {
	var n, c, d int
	fmt.Scanln(&n)
	fmt.Scanln(&c)
	fmt.Scanln(&d)

	result := findFirstNumber(n, c, d)
	if result != 0 {
		fmt.Println(result)
	} else {
		fmt.Println(" ")
	}
}
