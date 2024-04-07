package main

import "fmt"

func main() {
	var k, h, m int
	fmt.Scan(&k)

	h = k / 3600
	k = k - h
	m = k / 60 % 60

	fmt.Println("It is", h, "hours", m, "minutes")
}
