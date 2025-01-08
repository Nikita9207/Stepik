//Напишите "функцию голосования", возвращающую то значение (0 или 1),
//которое среди значений ее аргументов x, y, z встречается чаще.
//Входные данные
//Вводится 3 числа - x, y и z (x, y и z равны 0 или 1).
//Выходные данные
//Необходимо вернуть значение функции от x, y и z.

package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x)
	if x < 0 || x > 1 {
		fmt.Print("Ошибка ввода, введите число 0 или 1!")
		return
	}
	fmt.Scan(&y)
	if y < 0 || y > 1 {
		fmt.Print("Ошибка ввода, введите число 0 или 1!")
		return
	}
	fmt.Scan(&z)
	if z < 0 || z > 1 {
		fmt.Print("Ошибка ввода, введите число 0 или 1!")
		return
	}

	itog := vote(x, y, z)

	fmt.Print(itog)
}

func vote(x int, y int, z int) int {
	count0 := 0
	count1 := 0

	if x == 1 {
		count1++
	} else {
		count0++
	}

	if y == 1 {
		count1++
	} else {
		count0++
	}

	if z == 1 {
		count1++
	} else {
		count0++
	}

	if count1 > count0 {
		return 1
	} else {
		return 0
	}
}
