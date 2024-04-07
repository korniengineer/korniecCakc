package main

import "fmt"

func main() {
	var n, i, a int
	var slice []int
	fmt.Scan(&n)
	for ; i < n; i++ {
		fmt.Scan(&a)
		slice = append(slice, a)
	}
	fmt.Println(slice[3])
}
