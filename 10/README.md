## Canais

Forma de transmitir dados entre várias goroutines e lembrando que a func main também é uma goroutine. <br>
É regra que uma informação só pode ser criada por uma goroutine e só pode ser recebida por outra goroutine.
```
package main 

import (
    "fmt"
)

func main() {
    channel := make(chan int)
    
    go func() {
        channel <- 42
    }()
    
    fmt.Printf(<-channel)
}
```

Cada item no canal só pode ser lido uma vez. Quando uma goroutine está lendo do canal, ela espera algum valor vir.
```
package main

import (
	"fmt"
	"sync"
	"time"
)

var sg sync.WaitGroup

func consume(channel chan int) {
	fmt.Println(<-channel)
	fmt.Printf("Received!\n")
	sg.Done()
}

func send(channel chan int) {
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
```
Observe: a go function envia uma mensagem em 2 segundos. Por 2 segundos, a go function espera essa mensagem e, quando recebe, exibe.

### Canais direcionais

Basicamente, podemos receber nas funções:
```
// Apenas enviar
func send(channel chan<- int) { ... }

// Apenas receber
func receive(channel <-chan int) { ... }
```

### Range e Close

Podemos utilizar range assistindo um canal e dar um close quando mais nada for inserido.
```
func send(max int, channel chan<- int) {
    for i := 1; i <= max; i++ {
        channel <- i
    }
    close(channel)
}

func receive(channel <-chan int) {
    for value := range channel {
        fmt.Printf(value)
    }
}

func main() {
    channel := make(chan int)
    
    go send(10, channel)
    receive(channel)
}
```

O close basicamente é a função falando para fechar o canal e que mais nada será mandado. Isso evitará deadlock na função receive, impedindo que o for espere para sempre um retorno. <br>
Além disso, o comma ok também funciona aqui.
```
value, ok := <-channel
```
Ok dirá se value é um valor inexistente ou não.

### Convergência e divergência

É uma forma de ganhar eficiência. Basicamente, digamos que você tenha 5000 dados para processar. Aí você vai lá e fica chamando para cada dado, uma goroutine / worker que vai ler e enviar para um channel de exibir resultados. <br>
Uma estratégia é divergir primeiro e depois convergir. Funciona assim:
- Ao invés de ter só 1 worker / goroutine, você itera em um for para criar vários workers / goroutines
- O worker continua salvando em um outro channel que está exibindo em outra goroutine. <br> 

No primeiro passo, você divergiu os dados, fazendo o processamento ir n vezes mais rápido (em que n é o número de workers trabalhando no fall out). Ao fim, todos os workers mandam para o mesmo lugar (fall in).

### Context

Observe a utilização abaixo:
```
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {

		// recebe job
		case job, ok := <-jobs:
			if !ok {
				fmt.Println("worker", id, "finalizou")
				return
			}

			// simula processamento
			time.Sleep(50 * time.Millisecond)

			results <- job * 2

		// escuta cancelamento
		case <-ctx.Done():
			fmt.Println("worker", id, "cancelado")
			return
		}
	}
}

func main() {
	const totalJobs = 2000
	const totalWorkers = 3

	// timeout global de 3 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	jobs := make(chan int, totalJobs)
	results := make(chan int, totalJobs)

	var wg sync.WaitGroup

	// FAN-OUT → cria workers
	for i := 0; i < totalWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, results, &wg)
	}

	// envia jobs com respeito ao cancelamento
        loop:
        for i := 0; i < totalJobs; i++ {
            select {
            case jobs <- i:
        
            case <-ctx.Done():
                fmt.Println("parando envio de jobs")
                break loop
            }
        }

	close(jobs)

	// espera workers terminarem
	wg.Wait()
	close(results)

	// FAN-IN → lê resultados
	for r := range results {
		fmt.Println(r)
	}

	fmt.Println("processamento encerrado")
}
```
Cria um contexto com context.WithTimeout que irá emitir um cancel se passar de 3 segundos. Defer cancel para que no final da execução de main, tudo seja cancelado. <br>
Ao cancelar, ele envia para ctx o valor true para Done(). Ou seja, usamos select que o tempo todo verifica se há um Done() em toda goroutine a ser cancelada. Assim que cancel() vir, done é assionado e tudo para.
