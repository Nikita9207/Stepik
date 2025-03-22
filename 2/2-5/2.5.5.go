//Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку
//Sample Input:
//zaabcbd
//Sample Output:
//zcd

package main

import (
	"fmt"
	"strings"
)

func main() {
	var st string
	fmt.Scanln(&st)

	str := make(map[rune]int)

	for _, s := range st {
		str[s]++
	}

	result := strings.Builder{}

	for _, s := range st {
		if str[s] == 1 {
			result.WriteRune(s)
		}
	}

	results := result.String()

	fmt.Println(results)
}
