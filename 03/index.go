package main

import "fmt"

var slice []int = []int{1, 2, 3, 4, 5}

func main() {
	total := 0

	for _, value := range slice {
		total += value
	}

	if total > 0 {
		fmt.Println(total)
	} else {
		fmt.Println("Valor zero ou inferior")
	}

	x := []int{1, 2, 3, 4, 5}
	y := []int{6, 7, 8, 9, 10}

	// Isso é elemento dos índices 0,1 do x + elementos do índice 3, 4 do y (1, 2, 9, 10)
	x = append(x[:2], y[3:]...)

	fmt.Println(x)
}
