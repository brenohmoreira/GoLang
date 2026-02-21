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
    
        // b is a slice of bytes ([]byte)
	b, err := json.Marshal(group)

	if err != nil {
		fmt.Println(err)
	}

	os.Stdout.Write(b)
}

```

O os é biblioteca responsável por manipular arquivos. Ao passar o json para o 'os.Stdout.Write', escreve-se o json na tela. Devemos escrever essas informações em letras maiusculas, dado que tudo em Golang escrito em maiusculo é passível de ser importado. <br>
Isso significa que se você criar uma struct com nome maiusculo, você poderá acessar essa struct fora do arquivo também. <br>
Além disso, todos atributos da struct devem começar em maiuscula também, pois o json.Marshal não consegue enxergar as informações se não estiver.
Tudo em json em Golang é orientado a slices de byte []byte. Podemos converter para string passando por exemplo:
```
bString := string(b)
```

## Unmarshal

É o processo oposto, recebendo um slice de bytes (representando um json) e gerar um struct.
```
package main

import (
	"encoding/json"
	"fmt"
)

type Informacoes struct {
	Nome          string  `json:"Nome"`
	Sobrenome     string  `json:"Sobrenome"`
	Idade         int     `json:"Idade"`
	Profissao     string  `json:"Profissao"`
	Contabancaria float64 `json:"Contabancaria"`
}

func main() {
    // Usando `` para escapar as aspas. Transformando string json em slice of bytes
    sb := []byte(`{"Nome":"James","Sobrenome":"Bond","Idade":40,"Profissao":"Agente Secreto","Contabancaria":1000000.5}`)
    
    jamesbond := Informacoes{}
    
    err := json.Unmarshal(sb, &jamesbond)
    
    if err != nil {
        fmt.Println("error:", err)
    }
    
    fmt.Println(jamesbond)
    fmt.Println(jamesbond.Profissao)
}
```

Passamos um ponteiro de jamesbond para que a função possa alterar o valor na posição do objeto. <br>
Podemos utilizar Encoders para realizar esse processo sem alocar memória (ou seja, sem atribuir a uma slice of bytes intermediária antes de fazer a ação final).
O Encoder irá funcionar para qualquer coisa que implemente io.Writer (qualquer coisa que tenha o método Write([]byte)). Isso pode ser tanto o os.Stdout (perceba que ele pode usar .Write logo após), quanto para funções http.ResponseWriter que usaremos depois em APIs RESTful.
```
func main() {
    b, _ := json.Marshal(group)
    os.Stdout.Write(b)
    
    // Trocar para
    
    encoder := json.NewEncoder(os.Stdout)
    encoder.Encode(group)
    
    // Ou então
    
    json.NewEncoder(os.Stdout).Encode(group)
}
```