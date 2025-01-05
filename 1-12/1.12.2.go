//Напишите программу, принимающая на вход число N ( N ≥ 4 ), а затем N целых чисел-элементов среза.
//На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.

package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	b := make([]int, a, 10)

	for i := 0; i < a; i++ {
		fmt.Scan(&b[i])
	}
	fmt.Printf("%d", b[3])
}
