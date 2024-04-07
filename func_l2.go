package main

import "fmt"

func fibonacci(n int) int {
	var x1, x2, ans int
	x1 = 0
	x2 = 1
	for i := 1; i < n; i++ {
		ans = x1 + x2
		x1, x2 = x2, x2+x1
	}
	if n == 1 {
		return 1
	} else {
		return ans
	}
}

func main() {
	var a int
	a = fibonacci(5)
	fmt.Println(a)
}
