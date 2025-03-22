package main

import "fmt"

type Bike struct {
	On    bool
	Ammo  int
	Power int
}

func (t *Bike) RideBike() bool {
	if t.On && t.Power > 0 {
		t.Power--
		return true
	}
	return false
}

func (t *Bike) Shoot() bool {
	if t.On && t.Ammo > 0 {
		t.Ammo--
		return true
	}
	return false
}

func main() {

	//testStruct :=
	/*
	 * Экземпляр созданной вами структуры необходимо передать в качестве
	 * аргумента функции testStruct, которая выполнит проверку соблюдения
	 * всех условий задания/
	 */
	testStruct := &Bike{}
	fmt.Print(testStruct)
	//}
}
