## Concorrência x Paralelismo

Golang foi a primeira linguagem criada com multi-core em mente. Assim sendo, Golang torna fácil criar código concorrente.

Definição: padrão de design / paradigma (forma de organizar o código) em que várias funcionalidades acontecem ao mesmo tempo e de forma independente.
A concorrência é a preocupação principal que deve ser levada em conta por parte do desenvolvedor. É basicamente o design que vai permitir que o código faça várias coisas ao mesmo tempo, mas não simultaneamente. <be>
Mesmo com 1 núcleo, ainda assim a concorrência acontece. Não é sobre tarefas acontecendo literalmente ao mesmo tempo necessariamente. <br>
Por exemplo, um garçom serve 5 mesas (5 tarefas). Ele não fará isso ao mesmo tempo. Ele irá uma a uma (caso de 1 núcleo) e irá resolver. <br> 

Concorrência é o paradigma que permite que duas ou mais tarefas progridam juntas ao longo do tempo. Você basicamente desvincula a linearidade da execução do código. Enquanto isso, o paralelismo é a capacidade do hardware de fazer multiplas execuções dessas funções ao mesmo tempo, a depender do número de núcleos da GPU (em Golang, gerido automaticamente).

### Goroutines 

Utilizamos o package sync para criar goroutines. Para isso, definimos uma variável global que diz respeito a um grupo de goroutines:
```
var sg sync.waitGroup
```

Basicamente, devemos no inicio do método definir quantas goroutines haverão na execução daquele método. Para isso:
```
sg.Add(2)
```

Agora, devemos chamar nossas goroutines passando "go" no inicio da função:
```
go enumerateFirst()
go enumareteSecond()
```

Perfeito. Agora, antes de finalizar a execuçao de main, devemos dizer para não finalizar o processo enquanto todas as goroutines passadas para sg.Add sejam finalizadas. Então, antes de main terminar, passamos:
```
sg.Wait()
```

Mas o que define que a função foi finalizada? Para isso, no final de cada goroutine, passamos:
```
sg.Done() 
```

Um exemplo completo ficaria:
```
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
```

Atenção: falamos de duas goroutines que poderão sim ser executadas ao mesmo tempo. Agora, isso não significa que uma rodará primeiro que a outra. <br>
Basicamente, colocamos todas as goroutines a nível de runtime em um schedule. A segunda pode ser alocada para uma thread primeiro, por exemplo. As duas podem ter sido alocadas ao mesmo tempo, mas a execução de uma pode ter sido finalizada antes de GPU ter sido alocada para a segunda, por exemplo.