# Tipos Compostos

## Arrays 

Conjunto de x quantidade de valores de x tipo, não podendo variar. Um array de um tipo de y elementos nunca será igual a um array de mesmo tipo e com z elementos.
```
// x != y
var x [3]int
var y [4]int 

func main() {
    x[0] = 2
    y[0] = x[0]
    // len(x) e len(y) retornará o número de elementos destas arrays (3 e 4, respectivamente)
}
```

Para inicializar valores em uma array, podemos usar chaves:
```
var x = [3]int{1, 2, 3}
```
Geralmente, é usado somente slices. Arrays são usados para otimizar ao máximo no sentido de memória (desnecessário, quase sempre)

## Slices

Escolhemos um tipo e fazemos uma coleção de dados deste tipo. Podemos explicar slices pensando nelas como arrays com posições infinitas alocadas sob necessidade. São definidas de forma parecida, se diferenciando com a ausência de um valor limite de posições:
```
var x []int
```

Não podemos adicionar valores em slices utilizando indices pois não há nada garantido (não dará erro se o índice existir), a nível de compilador, que existirá esse índice. Por isso, utilizamos append:
```
var x = []int{1, 2, 3, 4}

x = append(x, 5)

// [1, 2, 3, 4, 5]
fmt.Printf(x) 
```

Para pegar partes da slice, temos:
```
var x []int = []int{1, 2, 3, 4}

// 2, 3, 4
var y []int = x[1:len(x)]
```

Para excluir itens de uma slide, combinamos o [x?:y?] com append + operador unfearl ...
```
var x []int = []int{1, 2, 3, 4, 5}
var y []int = []int{6, 7, 8, 9, 10}

// Isso é elemento dos índices 0,1 do x + elementos 8, 9, 10 do y
x = append(x[:2], y[3: ]...)
```

### Make

Utilizar slices é conveniente, pois a cada item a mais, o compilador apaga a array e cria uma nova para atender aquela necessidade. Isso, entretanto, tem um custo computacional. Para resolver esse problema, utiliza-se make.
O make pode ser definido:
```
make([]T, len, cap)
```
- O len é o length e diz respeito mais a slice
- A capacity diz mais respeito a array que a slice gera

Ao criar uma slice, por exemplo, passamos um valor inicial (assim como fazemos para a array) de posições.
Além disso, passamos uma capacidade. Ficaria assim:
```
slice := make([]int, 5, 10)
```

Isso vai criar, a grosso modo, uma array que estica. Ao invés de todo append criar uma array, como era feito com o slice antes, ele irá criar uma array de length definido no make e preencher todos os valores não passados com o valor zero do tipo.
```
// [0, 0, 0, 0, 0]
fmt.Printf(slice)
```

Não podemos dizer que:
```
slice[6] = 5
```
Isso dará erro, pois o length é 5. Porém, podemos utilizar append nele. 
Se todos os cinco espaços já estiverem preenchidos, ele irá "esticar" a array, permitindo que o tamanho se ajuste até atingir a capacidade.
Uma nova array só será criada quando a cap for atingida, onde cap é reajustado com o dobro da anterior.

### Array Subjacente

Quando se faz:
```
var x []int = []int{1, 2, 3, 4, 5}
var y []int = []int{6, 7, 8, 9, 10}

// Isso é elemento dos índices 0,1 do x + elementos 8, 9, 10 do y
x = append(x[:2], y[3: ]...)
```

O que é feito é que o slice x é reajustado e depois recortado. Ou seja, o x mudará.
Ou seja, isso nunca pode ser feito da seguinte maneira:
```
z = append(x[:2], y[3:]...)
```

Z irá mudar e X também. Sempre que utilizar append, a array do primeiro parâmetro deve ser a mesma onde está acontecendo a atribuição.

## Maps

O map se assemelha muito a um objeto em JavaScript. Basicamente, nele é utilizado o conjunto de índice/valor. 
Primeiro você define o tipo do índice e depois o tipo do valor. 

```
example := map[string]int
```

Isso pode ser definido com valores iniciais utilizando chaves:
```
example := map[string]int{
    "Breno": 123,
    "Honiele": 345
}
```

Para acessar, basta utilizar o índice e rapidamente o valor é encontrado.
```
// 123
fmt.Printf(example["Breno"])
```

Podemos definir valores depois utilizando:
```
example["Enzo"]: 678
```

As vezes, podemos acabar desejando um índice que não existe. Neste caso, o valor será o valor zero do tipo do map.
Porém, as vezes, teremos valores definidos que são os valores zero. Como podemos saber então se o valor obtido é um valor zero ou um valor definido com valor zero?
```
example["Ademir"] = 0

// ok comma
if teste, ok := example["Paulo"]; !ok {
    fmt.Printf("Paulo não existe")
} else {
    fmt.Printf("Paulo existe!")
}
```

### For range + map

Podemos utilizar o for range com map. O índice sera a key e o value será o valor respectivo daquela key.
```
example := map[string]int{
    "Breno": 123,
    "Honiele": 345
}

for key, value := range example {
    fmt.Printf(key, value)
}
```


Por fim, para apagar elementos do map. Fazemos:
```
// Passamos o map e depois o índice
delete(example, "Breno")
```