package main

import (
	"fmt"
)

var x int = 10

func main() {
	bytesNumber, _ := fmt.Println("Hello, world")
	fmt.Println(bytesNumber, x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T", float64(x))
}
