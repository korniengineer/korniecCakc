package main

import "fmt"

func main() {
	var num, sum int
	fmt.Scan(&num)
	sum = (num-1)%9 + 1
	fmt.Println(sum)
}
