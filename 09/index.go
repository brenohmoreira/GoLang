package main

import (
	"fmt"
	"runtime"
	"sync"
)

var gr sync.WaitGroup
var mutex sync.Mutex

func main() {
	goRoutines := 1000
	contador := 0

	gr.Add(goRoutines)

	for i := 0; i < goRoutines; i++ {
		go func() {
			mutex.Lock()

			valor := contador
			runtime.Gosched()
			contador = valor + 1

			mutex.Unlock()

			gr.Done()
		}()
	}

	gr.Wait()

	fmt.Printf("Total contador: %d\n", contador)
}
