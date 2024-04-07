package main

import "fmt"

func minimumFromFour() int {
	var n, min int
	fmt.Scan(&min)
	for i := 1; i < 4; i++ {
		fmt.Scan(&n)
		if n < min {
			min = n
		}
	}
	return min
}

func main() {
	var minNumber int
	minNumber = minimumFromFour()
	fmt.Println(minNumber)
}
