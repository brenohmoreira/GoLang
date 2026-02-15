## Structs

Structs são tipos que podem armazenar dados com tipos diferentes. É uma estrutura variada.
Ele é definido em global scope e utilizado como um tipo. Por isso, sua definição necessita ser dar seguinte forma:
```
type car struct {
    id int
    marca string
    modelo string
    cor string
    estaLigado bool
}
```

Para utilizarmos a partir disso, segue:
```
func main() {
    carro := car{
        id: 0,
        marca: "Marca não sei",
        modelo: "Modelo não sei",
        cor: "Amarelo rosado",
        estaLigado: false
    }
    
    carroNaoDefinido := car{}
    
    fmt.Println(carro.modelo)
    
    // Retornará os valores zero: 0 "" "" """ false
    fmt.Println(carroNaoDefinido)
}
```

## Funções

Funções são utilizadas para abstrair funcionalidades e reutilizar códigos. O modelo é:
```
func (receiver) nome(...parameters) type {
    // Instruction here
}
```
Sendo que receiver é opcional e será abordado depois. Para chamar, usamos:
```
nome(...arguments)
```
Type pode ser o tipo do retorno ou nada, se for void. Podemos também fazer uma função que retorna vários valores. O ...T é o parâmetro variádico, significa que podem receber n valores int (no caso de ...int). O elemento variádico sempre deve ser o último definido.
```
func soma(x ...int) (int, int) {
    soma := 0
    
    for _, value := range x {
        soma += value
    }
    
    return soma, len(x) 
}

func main() {
    soma, tamanho := soma(1, 2, 3)
    
    // Retornará 6, 3
    fmt.Printf(soma, tamanho)
}
```
Internamente, o ...int faz x virar uma slice internamente.

