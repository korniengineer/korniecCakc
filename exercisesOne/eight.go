package main

import (
	"fmt"
)

func main() {
	var n, min, count, j int
	fmt.Scan(&n)
	min = n
	for i := 0; i < n; i++ {
		fmt.Scan(&j)
		if j < min {
			min = j
			count = 1
		} else if j == min {
			count++
		}
	}
	fmt.Println(count)
}
