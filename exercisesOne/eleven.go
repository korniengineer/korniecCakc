package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	if num == 1 {
		fmt.Println("1 korova")
	} else if num < 5 && num > 1 {
		fmt.Println(num, "korovy")
	} else if num > 4 && num < 21 {
		fmt.Println(num, "korov")
	} else if num > 20 && num%10 == 1 {
		fmt.Println(num, "korova")
	} else if num > 21 && num%10 < 5 {
		fmt.Println(num, "korovy")
	} else {
		fmt.Println(num, "korov")
	}
}
