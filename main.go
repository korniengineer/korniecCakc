package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func res(a int, b int) int {
	return a - b
}

func umn(c, d int) int {
	return c * d
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool
var ab, ac = 1, 2

func main() {
	fmt.Println("This is IT!")
	fmt.Println(add(15, 20))
	fmt.Println(res(20, 17))
	fmt.Println(umn(3, 5))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
	var i int
	fmt.Println(i, c, python, java)
	var ad, ae, af = true, false, "no!"
	fmt.Println(ab, ac, ad, ae, af)
	k := 3
	fmt.Println(k)
}
