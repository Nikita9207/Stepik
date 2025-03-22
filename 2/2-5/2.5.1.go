//На вход подается строка. Нужно определить, является ли она правильной или нет.
//Правильная строка начинается с заглавной буквы и заканчивается точкой.
//Если строка правильная - вывести Right иначе - вывести Wrong

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var str string

	scan := bufio.NewReader(os.Stdin)
	str, _ = scan.ReadString('\n')
	strS := strings.TrimSpace(str)

	if examination(strS) {
		fmt.Println("Right")
	} else {
		fmt.Print("Wrong")
	}

}

func examination(strS string) bool {
	runes := []rune(strS)
	if len(strS) > 0 && unicode.IsUpper(runes[0]) && strings.HasSuffix(strS, ".") {
		return true
	}
	return false
}
