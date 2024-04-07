package main

import "fmt"

func main() {
	var a, b, c, sqr1, sqr2 int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	sqr1 = c * c
	sqr2 = (a * a) + (b * b)

	if sqr1 == sqr2 {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}
}
