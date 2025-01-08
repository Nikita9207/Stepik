//Напишите функцию f(), которая будет принимать строку text и выводить ее (печатать на экране).
//От вас требуется дописать только эту функцию, считайте что функция main() уже объявлена, считывать с консоли ничего не нужно!
//Пакет "fmt" уже импортирован!

package main

import "fmt"

func main() {
	var text = "Hello!"
	text = f(text)
	fmt.Print(text)
}
func f(text string) string {
	fmt.Print(text)
	var a = text
	return a
}
