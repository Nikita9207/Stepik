//На вход подаются a и b - катеты прямоугольного треугольника. Нужно найти длину гипотенузы
//Sample Input:
//6 8
//Sample Output:
//10

package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64

	_, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println("Ошибка")
	}

	d := a*a + b*b

	c := math.Sqrt(d)

	fmt.Println(c)
}
