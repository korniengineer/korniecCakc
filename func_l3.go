package main

import "fmt"

func sumInt(a ...int) (int, int) {
	var i int
	sum := 0
	for _, elem := range a {
		i++
		sum = sum + elem
	}
	return i, sum
}

func main() {
	var n, m int
	n, m = sumInt(3, 5)
	fmt.Println(n, m)
}
