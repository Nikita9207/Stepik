//Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.
//Входные данные
//Программа получает на вход два числа. Гарантируется, что цифры в числах не повторяются. Числа в пределах от 0 до 10000.
//Выходные данные
//Программа должна вывести цифры, которые имеются в обоих числах, через пробел. Цифры выводятся в порядке их нахождения в первом числе.
//Внимание! Если вам сложно решить эту задачу, пропустите и изучайте материал дальше, затем вернетесь к этой задаче позже.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a, b int
	var num1 []int
	var num2 []int
	var result []int
	fmt.Scan(&a, &b)

	for a > 0 {
		aNum := a % 10
		num1 = append(num1, aNum)
		a -= aNum
		a /= 10
	}

	for b > 0 {
		bNum := b % 10
		num2 = append(num2, bNum)
		b -= bNum
		b /= 10
	}

	for _, v1 := range num1 {
		for _, v2 := range num2 {
			if v1 == v2 {
				result = append(result, v1)
			}
		}
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	var sresult []string

	for _, v := range result {
		sresult = append(sresult, strconv.Itoa(v))
	}

	fmt.Print(strings.Join(sresult, " "))
}
