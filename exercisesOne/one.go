package main

import "fmt"

func main() {
	var num, a, b, c int
	fmt.Scan(&num)
	a = num % 10
	b = num % 100 / 10
	c = num % 1000 / 100
	fmt.Print(a)
	fmt.Print(b)
	fmt.Print(c)

}
