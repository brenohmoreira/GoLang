package main

import (
	"fmt"
	"sync"
	"time"
)

var sg sync.WaitGroup

func consume(channel <-chan int) {
	fmt.Println(<-channel)
	fmt.Printf("Received!\n")
	sg.Done()
}

func send(channel chan<- int) {
	fmt.Printf("Sending...\n")
	time.Sleep(2 * time.Second)
	fmt.Printf("Sended...\n")

	channel <- 42
	sg.Done()
}

func main() {
	channel := make(chan int)

	sg.Add(2)

	go send(channel)
	go consume(channel)

	sg.Wait()
}
