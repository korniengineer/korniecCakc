package main

import "fmt"

func main() {
	var workArray [10]int
	var num, ind, indF, indS int
	for i := 0; i < len(workArray); i++ {
		fmt.Scan(&num)
		workArray[i] = num
	}

	// Создание массива с вводом данных

	for i := 0; i < 3; i++ {
		fmt.Scan(&ind)
		indF = ind
		fmt.Scan(&ind)
		indS = ind
		ind = workArray[indF]
		workArray[indF] = workArray[indS]
		workArray[indS] = ind

	}
	for _, elem := range workArray {
		fmt.Print(elem)
	}
	fmt.Print("type ok")
}
