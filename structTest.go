package main

import "fmt"

func main() {
	testStruct := mainStruct{true, 12, 3}
	fmt.Println(testStruct.Shoot())
}

type mainStruct struct {
	On          bool
	Ammo, Power int
}

func (m *mainStruct) Shoot() bool {
	if m.On == false || m.Ammo < 1 {
		return false
	} else {
		m.Ammo--
		return true
	}
}

func (m *mainStruct) RideBike() bool {
	if m.On == false || m.Power < 1 {
		return false
	} else {
		m.Power--
		return true
	}
}
