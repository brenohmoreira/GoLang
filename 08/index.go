package main

import (
	"fmt"
	"sync"
)

func enumerateFirst() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("First: %d\n", i)
	}

	sg.Done()
}

func enumareteSecond() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("Second: %d\n", i)
	}

	sg.Done()
}

var sg sync.WaitGroup

func main() {
	sg.Add(2)

	go enumerateFirst()
	go enumareteSecond()

	sg.Wait()
}
