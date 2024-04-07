package main

import "fmt"

func main() {
	var a []int
	var b []int = []int{1, 2, 3}
	c := []int{1, 2, 3}
	d := []int{1: 12, 2: 35}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	slice := make([]int, 10, 10)
	fmt.Println(slice)

	initialUsers := [8]string{"Bob", "Alice", "Margo", "John", "James", "William", "Sarah", "Sergo"}
	users_1 := initialUsers[2:6]
	users_2 := initialUsers[:4]
	users_3 := initialUsers[3:]
	fmt.Println(users_1)
	fmt.Println(users_2)
	fmt.Println(users_3)

	ab := []int{1, 2, 3}
	ab = append(ab, 4, 5)
	fmt.Println(ab)
}
