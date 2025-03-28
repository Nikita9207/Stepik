//Внимательно прочитайте ВСЕ условия и подсказки чтобы правильно решить задачу!
//
//Требуется вычислить период колебаний (t) математического маятника (мы округлили некоторые значения для удобства проверки),
//для этого нужно найти циклическую частоту колебания пружинного маятника (w), в формуле w встречается масса которую также нужно найти,
//все нужные формулы приведены ниже:
//Напишите три функции, каждая из которых будет выполнять конкретную формулу. Название функций обязательно должны
//соответствовать букве формулы: T(), W() и M(). Для того чтобы найти t - необходимо сначала найти w, и т.д.
//Так что используйте результат функции W() в формуле функции T() - то-есть вызывайте функцию W() в T(). Аналогично и с W(), M().

//t = 6 / w , w =2| k / m , m = p ∗ v
//ВАЖНО! Считайте, что пакет main уже объявлен, а также функция main() с вызовом ВАШЕЙ будущей функции T() уже есть.
//Несмотря на то, что тестирование будет через ввод-вывод, вам НЕ требуется вводить и выводить что-либо.
//Для подсчета используйте УЖЕ ВВЕДЕННЫЕ ГЛОБАЛЬНЫЕ переменные k,p,v ТИПА float64 !!!
//
//Пакет math уже импортирован! Напоминаю: корень (sqrt) можно найти с помощью пакета "math", например:
//fmt.Println(math.Sqrt(9))
// результат: 3
//
//Sample Input:
//1296 6 6
//Sample Output:
//1

package main

import (
	"fmt"
	"math"
)

var k, p, v float64

func main() {
	_, err := fmt.Scanln(&k, &p, &v)
	if err != nil {
		fmt.Println("Ошибка")
	}

	t := T()
	fmt.Printf("T = %.2f", t)
}

func M() float64 {
	m := p * v
	return m
}

func W() float64 {
	m := M()
	d := k / m
	w := math.Sqrt(d)
	return w
}

func T() float64 {
	w := W()
	t := 6 / w
	return t
}
