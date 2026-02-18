## JSON 

A biblioteca para manipular JSON é: "encoding/json" <br>
Ao passar um type customizado com struct para a função json.Marshal, você irá converter aquele objeto em um JSON. <br>
O método json.Marshal retorna o próprio json e o erro logo após, que, se recebido, pode ser tratado da forma que preferir. Segue abaixo um modelo:
```
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

func main() {
	group := ColorGroup{ID: 1, Name: "Red", Colors: []string{"Blue", "Green"}}

	b, err := json.Marshal(group)

	if err != nil {
		fmt.Println(err)
	}

	os.Stdout.Write(b)
}

```

O os é biblioteca responsável por manipular arquivos. Ao passar o json para o 'os.Stdout.Write', escreve-se o json na tela.