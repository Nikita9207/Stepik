//Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования.
//Длина пароля должна быть не менее 5 символов, он может содержать только арабские цифры и буквы латинского алфавита.
//На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"
//
//Sample Input:
//fdsghdfgjsdDD1
//Sample Output:
//Ok

package main

import (
	"fmt"
	"unicode"
)

func valid(password string) bool {
	if len(password) < 5 {
		return false
	}

	for _, st := range password {
		if !unicode.IsLetter(st) && !unicode.IsDigit(st) {
			return false
		}
		if unicode.IsLetter(st) && !unicode.Is(unicode.Latin, st) {
			return false
		}
	}
	return true
}

func main() {
	var password string
	fmt.Scanln(&password)

	if valid(password) {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}
