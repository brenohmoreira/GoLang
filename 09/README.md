## Condição de Corrida

Digamos que uma goroutine seja chamada 10 vezes. Para cada vez, ele captura o valor, dorme por 1-2 segundos e depois increnta esse valor, alterando a variável de escopo global. Ao final da execução, queremos o resultado total.
Qual será o resultado? <br>
Se a resposta foi, dez, está errado. O ponto é que todas elas pegam um valor dessa variável a 2 segundos, dormem e depois atribuem. Mas durante esses dois segundos, pode ser que a variável tenha mudado por causa das outras goroutines. Nesse caso, não há valor correto a ser esperado, tudo que sabemos com certeza é que será maior do que o valor inicial da variável de global scope e menor do que o número de repetições das goroutines. <br>
Existem muitas formas de se resolver isso, para começar, faremos com mutex. Ante disso, segue o problema exemplo:
```
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var gr sync.WaitGroup

func main() {
	goRoutines := 1000
	contador := 0

	gr.Add(goRoutines)

	for i := 0; i < goRoutines; i++ {
		go func() {
			valor := contador
			runtime.Gosched()
			contador = valor + 1
			gr.Done()
		}()
	}

	gr.Wait()

	fmt.Printf("Total contador: %d\n", contador)
}
```
Isso, que era para dar 1000, dá muito menos. Vamos resolver isso agora.

### Mutex 

O mutex, por meio do Lock e do Unlock, cria uma região crítica onde cada goroutine só pode executar ela uma por vez. As outras esperam.
```
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
```