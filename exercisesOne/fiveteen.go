package main

import "fmt"

func main() {
	var n, k, g int
	var array []int
	fmt.Scan(&n)
	fmt.Scan(&g)
	for ; n >= 0; n = n / 10 {
		k = n % 10
		if k == g {
			continue
		}
		array = append(array, k)
	}
	fmt.Println(array)
}
