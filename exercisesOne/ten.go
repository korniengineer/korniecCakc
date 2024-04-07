package main

import "fmt"

func main() {
	var a, b, max int
	max = -1
	fmt.Scan(&a)
	fmt.Scan(&b)
	for ; b >= a; b-- {
		if b%7 == 0 {
			max = b
			break
		}
	}
	if max == -1 {
		fmt.Println("NO")
	} else {
		fmt.Println(max)
	}
}
