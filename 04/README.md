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
Internamente, o ...int faz x virar uma slice internamente. Quando usamos parâmetros variádicos, é bom ressaltar que a falta de envio não resulta em problema nenhum, pois será traduzido para uma slice vazia.

### Métodos 

Métodos fazem referência a dar um contexto de um tipo para uma função passando esse tipo no receiver.
Ao criar um receiver na função e passar um tipo x, significa que uma variável de tipo x pode implementar essa função e durante a execução dessa variável, seu contexto / informações serão armazenados no próprio receiver.

```
type car struct {
    id int 
    cor string 
}

func (c car) printCar(name string) {
    fmt.Printf("ID: %d\nNOME: %s\nCOR: %s", c.id, c.cor, name)
}

func main() {
   mercedes := car{ id: 0, cor: "Branco" } 
   mercedes.printCar("Mercedes") 
}
```
Função printCar está anexada ao tipo car, ou seja, não pode ser utilizada sem tipo ou com outros tipos.
Criar uma variável de uma type struct neste contexto é o mesmo que um objeto de POO.

### Interfaces e polimorfismo
Ao criar uma interface, podemos passar quais são as implementações necessárias para chamar essa interface.
É totalmente implicito, não ha diretivas como "implements" do Java. 
Se uma interface "Animais" exige implementação do método "fazerBarulho", todos tipos (ex: cachorro, gato) que tenham um método "fazerBarulho" automaticamente implementam a interface "Animais".
```
type gato struct {
    nome string 
}

func (g gato) fazerBarulho() {
    fmt.Println("Miau")
}

type cachorro struct {
    nome string
}

func (c cachorro) fazerBarulho() {
    fmt.Println("Auau")
}

type animal interface {
    fazerBarulho()
}
```
Pronto, agora todos que tenham o método "fazerBarulho" implementam "animal". 
O que podemos fazer agora é chamar fazerBarulho. Para isso, primeiro vamos criar um cachorro e um gato.
```
fuc main() {
    cat := gato{ nome: "Xodó" }
    dog := cachorro{ nome: "Rex" }
}
```

Gato é gato, mas também é animal. Cachorro é cachorro, mas também é animal. Podemos criar uma função intermediária agora
que recebe como parâmetro um animal e chama o seu método fazerBarulho.
```
func efetuarBarulho(a animal) string {
    // Isso é polimorfismo!
    return a.fazerBarulho() 
}   

fuc main() {
    cat := gato{ nome: "Xodó" }
    dog := cachorro{ nome: "Rex" }
    
    // Miau
    fmt.Printf(efetuarBarulho(cat))
        
    // Auau
    fmt.Printf(efetuarBarulho(dog))
}
```
Um mesmo método realizando ações diferentes. Isso é polimorfismo. A ação vai depender do tipo em runtime.

### Callback

É quando é passado uma function como parâmetro para outra function a fim de ser utilizado dentro da função principal para obtenção de um resultado.
Ao passar uma função como callback, não utilizemos os parênteses, pois não estamos a realizar uma chamada, mas sim dando a uma função a aplicabilidade de outra função.
```
func soma(x ...int) int {
    soma := 0
    for _, value := range x {
        soma += value 
    }
    return soma 
}

func somaSomentePares(sumFunction func(x ...int) int, y ...int) int {
    somaSlice := []int
    for _, value := range y {
        if value % 2 == 0 {
            somaSlice = append(somaSlice, value)
        }
    }
    somaPares := sumFunction(somaSlice...)
    return somaPares
} 

func main() {
    // Somará 2, 4 
    t := somaSomentePares(soma, []int{1, 2, 3, 4, 5})
    
    // 6
    fmt.Printf(t)
}
```

### Closures

Utilizado quando uma função retorna uma função e este retorno é armazenado em uma variável.
Ao fazer isso, o contexto da função fica salvo e existente enquanto a variável existir.
Isso significa que, se criarmos:
```
func closureContext() func () int {
    x := 0
    return func () int {
        return ++x 
    }
}

func main() {
    // Aqui x := 0 foi criado no contexto e, para A, ele existirá
    a := closureContext()
        
    // 1
    a() 
    // 2
    a() 
    // 3
    a()
}
```


### Defer 

Quando usamos um statement com defer, basicamente dizemos que aquele statement deve ser deixado por último.
A primeira instrução com defer criada é a última a ser executada no fim da função.
```
func main() {
    /*
    Aparecerá: 
    3
    4
    2
    1
    */
    
    defer fmt.Printf("1")
    defer fmt.Printf("2")
    fmt.Printf("3")
    fmt.Printf("4")
}
```
O defer é ótimo para rotinas que devem fazer algo no final com algo relacionado ao início. Exemplo: abre conexão com o banco e cria um defer para fechar a conexão logo em seguida.