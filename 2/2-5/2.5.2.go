//На вход подается строка. Нужно определить, является ли она палиндромом. Если строка является палиндромом - вывести Палиндром
//иначе - вывести Нет. Все входные строчки в нижнем регистре.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var str string
	scan := bufio.NewReader(os.Stdin)
	str, _ = scan.ReadString('\n')
	str = strings.ToLower(str)
	str = strings.TrimSpace(str)
	//fmt.Println(str)

	if isPalindrome(str) {
		fmt.Print("Палиндром")
	} else {
		fmt.Print("Нет")
	}
}

func isPalindrome(str string) bool {
	var pal string
	for _, s := range str {
		pal = string(s) + pal
	}
	if pal == str {
		return true
	} else {
		return false
	}
}
