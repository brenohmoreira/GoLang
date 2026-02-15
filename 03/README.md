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

