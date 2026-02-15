package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello World\n")

	for horas := 1; horas <= 24; horas++ {
		fmt.Printf("Horas: %d\n", horas)

		for minutes := 1; minutes <= 60; minutes++ {
			fmt.Printf("%02d:%02d\n", horas, minutes)
		}

		fmt.Println("\n")
	}
}
