package main

import "fmt"

func main() {
	var n, n1, c int
	var slice []int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&n1)
		slice = append(slice, n1)
	}
	for i := 0; i < len(slice); i++ {
		if slice[i] > 0 {
			c++
		}
	}
	fmt.Println(c)

}
