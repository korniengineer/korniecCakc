package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	if a+b > c || a+c > b || b+c > a {
		fmt.Println("Существует")
	} else {
		fmt.Println("Не существует")
	}
}
