//На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)

package main

import "fmt"

func main() {
	var str string
	fmt.Scanln(&str)

	result := ""
	for i := 1; i < len(str); i += 2 {
		result += string(str[i])
	}

	fmt.Println(result)
}
