Para rodar uma aplicação em go basta no diretório principal do projeto utilizar:
    ```
    go run .
    go run [file]
    ```

- Variáveis definidas como "_" não devem ser utilizada (como dizer que a variável é um retorno inutil. É útil para funções que retornam dois ou valores, mas só precisamos de um)
- Todos os tipos em go implementam uma interface {} (interface vazia) 
- Uma expressão não gera efeitos colaterais e pode atribuir algo, ex: 2 + 5
- Uma declaração gera efeitos colaterais, são usados em ifs/while/do...while, etc
- O arquivo inicial executado em um projeto é aquele que importa o package main
    - A func main é a primeira a ser executada e última a ser finalizada

## Tipos

Antes de inicializar uma variável, ela, ao ser declarada, possui um valor zero que depende do tipo.

- var y string => ""
- var y int => 0
- var y float32 => 0.0
- var y bool => false
- pointers, functions, interfaces, slices, channels, maps => nill

Há também o tipo const. São valores imutáveis utilizado em global scope. Se não houver valor definido, seu tipo só sera definido na hora que for utilizado.
```
const x = 10
var y int
var z float32 

func main() {
    y = x
    z = x
    
    // Espelhou x, mas é int pois var y foi definida como inteira 
    fmt.Printf("%T\n", y)
       
    // Espelhou x, mas é float32 pois var z foi definida como ponto flutuante de 32 bits
    fmt.Printf("%T\n", z)
}
```

### Primitivos 

Os tipos primitivos da Golang é:
- Number (Int/Float32/Float64)
- String
    - Pode ser Literal Strings ou Raw Literal Strings. As Literal Strings são definidas com aspas duplas ("") e interpretadas, fazendo com que escapes, quebra de linha e paragrafos sejam possíveis. A Raw Literal Strings são definidos com acentos graves (``) e estas são literalmente todas transformadas em texto.
- Bool

Se você utilizar :=, será realizada uma tipagem automática. É utilizado para inicializar variáveis (além de também atribuir). Não pode ser usada em global scope.
Se você utilizar =, será realizada uma atribuição. Após uma variável estar inicializa, é possível atribuir valores a ela com o equals. Em global scope, utiliza-se para inicializar também (com var).

Tipos em Go são estáticos, ou seja, não é possível alterar o tipo de uma variável ao decorrer de uma chamada. Para definir um tipo, apenas damos espaços e colocamos.
Só podemos dizer explicitamente o tipo se for var.

Exemplo: 
```
var x int = 10
```

### Compostos

Os tipos compostos são: slice, array, struct, map
O ato de de definir, criar e estruturar tipos compostos chama-se de composição.


### Tipos personalizados

type customType int 
var customVariable customType = 3

Custom type tem um tipo adjacente int

### Conversion 

Go só há conversion, nunca casting. Para isso, basta utilizar t(x) onde t é o tipo desejado e x é a variável a ser convertida.
Ex: 
intVariable := 3
floatVariable := float(intVariable)

## Bit 

Um bit pode ser 0 ou 1. Ou seja, cada bit assume um valor binário e um conjunto de n bites tem 2^n possíveis combinações de valores. 
Por exemplo, 8 bits tem 256 possibilidades (0 - 255). 

### Bits aplicados  

Segue o range de um uint: 
- uint8 -> 0 - 255 (8 bits, 2^8)
- uint16 -> 0 - 65535 (16 bits, 2^16)
- uint32 -> 0 - (2^32 - 1) (32 bits, 2^32)
- uint64 -> 0 - (2^64 - 1) (64 bits, 2^64)

O int padrão utiliza um bit para dizer se o valor é negativo ou positivo. Ou seja, 7 bits para definição do range e 1 bit (0 ou 1) para o sinal(- ou +).

Segue o range de um int:
- int8 -> -128 to 127 (7 bits, 2^7 e 1 bit para definir -128 to -1 ou 0 to 127 = 8 bits)
- int16 -> -32768 to 32767 (15 bits, 2^15 e 1 bit para sinal = 16 bits)
- int32 -> -(2^32) to (2^31 - 1) (31 bits, 2^31 e 1 bit para sinal = 32 bits) 
- int64 -> -(2^63) to (2^63 - 1) (63 bits, 2^63 e 1 bit para sinal = 64 bits)

Overflow acontece quando uma variável definida com um valor limitado (ex: int8), a variável se encontra com o valor máximo definido para int8 e implicitamente algo acresce nessa variável (ex: variavel++). Isso faz ele voltar ao limite menor, ou seja, no caso de int8, vai de 255 para 0.

O tipo byte é o valor no range inteiro que aquele caractere ocupa. Se você criar uma string e converter para um array de bytes, você chegará na array de números do range que formam aquela string.
```
text := "Hello, World" 
bytesText := []byte(text)
```

Cada caracter pode ter até 4 valores uint8, no caso de utf8, variando de caracter para caracter.
Se um byte tem 8 bits e o utf-8 pode ser representado até com 4 bytes, então temos que cada byte do utf-8 que virá na array []byte vai de 0-255 (representação decimal).

## Iota

Ao definir uma const como iota, você dá a ele valores untyped inteiros sucessivos. 

```
const {
    x = iota,
    y = iota,
    z = iota,
    _ = iota,
    b = iota
}
```
x será 0, y será 1, z será 2, b será 4