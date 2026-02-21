## Packages

### Sort

Biblioteca de ordenação. Podemos passar para ele um slice e ele irá ordenar de acordo com o tipo do slice.
```
import (
    "fmt",
    "sort"
)

func main() {
    letters := []string{"A", "B", "D", "C", "Z", "X"}
    sort.Strings(letters)
    
    // [A, B, C, D, X, Z]
    fmt.Printf(letters)
}
```
Há também:
```
sort.Floats(x)
sort.Ints(x)
sort.Strings(x)
```

#### Sort customizado

Para types structs, precisamos cusstomizar o sort:
```
package main

import (
	"fmt"
	"sort"
)

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	pessoas := []Pessoa{
		{"Ana", 25},
		{"Carlos", 20},
		{"Bruno", 30},
	}

	// Ordenar por idade (crescente)
	sort.Slice(pessoas, func(i, j int) bool {
		return pessoas[i].Idade < pessoas[j].Idade
	})

	fmt.Println(pessoas)
}
```
É possível fazer pelo método de interfaces também. Basicamente, criamos um tipo dque usaremos como troca e esse tipo deve implementar 3 métodos: Len, Less e Swap. <br>
Basicamente, ele irá perguntar o tamanho com Len, depois verificar com Less e se necessário, trocar com Swap. Tudo isso implicitamente.
```
type PorIdade []Pessoa 

func (p PorIdade) Len() int { return len(p) }
func (p PorIdade) Less(i, j int) bool { return p[i].Idade < p[j].Idade }
func (p PorIdade) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func main() {
	pessoas := []Pessoa{
		{"Ana", 25},
		{"Carlos", 20},
		{"Bruno", 30},
	}

	// Ordenar por idade (crescente)
	sort.Sort(PorIdade(pessoas))
}
```

## Bcrypt

Package que permite a criptografia do tipo hash, utilizado geralmente em senhas ou dados sensíveis no geral. <br>
Para instalar a biblioteca, temos que, além do import, fazer:
```
go get -u golang.org/x/crypto/bcrypt
```
Isso irá criar o require correto em go.mod <br> 
Para utilizar, usamos:
```
password := "zanelor12"

hash, err := bcrypt.GenerateFromPassword([]byte(password), 4)

if err != nil {
    fmt.Println(err)
}

fmt.Println(string(hash))
```
O número 4 é o custo computacional. Quanto maior, mais seguro o hash é, mais custoso para gerar também. <br>
Para comparar senhas, por exemplo, não é possível retornar o hash para a forma de senha e nem recomendável (mesmo se fosse possível). Para isso, usamos o método para comparar.
```
err := bcrypt.CompareHashAndPassword(hash, []byte("zanelsor12"))

if err == nil {
    fmt.Printf("Senha correta")
} else {
    fmt.Printf("Senha incorreta")
}
```
Se não for o mesmo, ele irá retornar nil. Caso contrário, irá retornar um erro.