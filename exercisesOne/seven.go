package main

import "fmt"

func main() {
	var n, k, j int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&j)
		if j == 0 {
			k++
		}
	}
	fmt.Println(k)
}
