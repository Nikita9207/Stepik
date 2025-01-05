//По данному трехзначному числу определите, все ли его цифры различны.
//Формат входных данных
//На вход подается одно натуральное трехзначное число.
//Формат выходных данных
//Выведите "YES", если все цифры числа различны, в противном случае - "NO".

package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scanln(&number)

	if number < 100 || number > 999 {
		fmt.Println("Неверный ввод: должно быть трехзначное число.")
		return
	}

	digit1 := number / 100
	digit2 := (number / 10) % 10
	digit3 := number % 10

	if digit1 != digit2 && digit2 != digit3 && digit1 != digit3 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
