//Даются две строки X и S. Нужно найти и вывести первое вхождение подстроки S в строке X. Если подстроки S нет в строке X - вывести -1

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, x string
	scan := bufio.NewReader(os.Stdin)
	s, _ = scan.ReadString('\n')
	x, _ = scan.ReadString('\n')
	s = strings.TrimSpace(s)
	x = strings.TrimSpace(x)

	fmt.Println(strings.Index(s, x))
}
