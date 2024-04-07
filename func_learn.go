package main

import "fmt"

func hello() {
	fmt.Println("Hello World!")
}

func add(x int, y int) {
	z := x + y
	fmt.Println("x + y = ", z)
}

func add_2(x, y int, a, b, c float32) {
	z := x + y
	d := a + b + c
	fmt.Println("x + y = ", z)
	fmt.Println("a + b + c = ", d)
}

func main() {
	hello()
	add(10, 12)
	add_2(2, 3, 5.1, 6.8, 3.1)
}
