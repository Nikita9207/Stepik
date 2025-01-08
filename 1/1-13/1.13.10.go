//Самое большое число, кратное 7
//Найдите самое большее число на отрезке от a до b, кратное 7 .
//Входные данные
//Вводится два целых числа a и b (a≤b).
//Выходные данные
//Найдите самое большее число на отрезке от a до b (отрезок включает в себя числа a и b), кратное 7 , или выведите "NO" - если таковых нет.

package main

import (
	"fmt"
)

func main() {
	var a, b int

	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	max := -1

	for i := b; i >= a; i-- {
		if i%7 == 0 {
			max = i
			break
		}
	}

	if max == -1 {
		fmt.Println("NO")
	} else {
		fmt.Println(max)
	}
}
