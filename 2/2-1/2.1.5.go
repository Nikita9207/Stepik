//Напишите функцию sumInt, принимающую переменное количество аргументов типа int, и возвращающую количество полученных
//функцией аргументов и их сумму. Пакет "fmt" уже импортирован, функция и пакет main объявлены.
//Пример вызова вашей функции:
//a, b := sumInt(1, 0)
//fmt.Println(a, b)
//Результат: 2, 1

package main

import "fmt"

func main() {
	a, b := sumInt(1, 0)
	fmt.Println(a, b)
}

func sumInt(arg ...int) (a int, b int) {
	b = 0
	for _, a := range arg {
		b += a
	}

	return len(arg), b
}
