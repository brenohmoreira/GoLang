# Loops e condicionais

A grande diferença de loops e condicionais da Golang para outras linguagens, é que o parentêses é omitido em casos que não há inversão de valores com !

## Loops

Não existe while ou do...while em Golang. Para realizar loops, é utilizado apenas o for. Segue abaixo a forma mais rudimentar de fazê-lo:
```
for x := 0; x < 10; x++ {
    // Instruction here
}
```

A inicialização e o incremento/decremento podem ser tirados. O bloco abaixo funciona igual ao bloco acima.
```
x := 0
for x < 10 {
    // Instruction here
    x++
}
```

Há também o que seia o foreach da Golang:
```
var x []int{1, 2, 3}

for i, value range x {
    // Retornará o item da iteração e o índice do item
    fmt.Printf(i, value)
}
```

O statement break encerra o loop. O statement continue, por outro lado, força a ir para a próxima iteração.

## Condicionais

Para realizar uma condicional, segue-se o padrão do bloco abaixo:
```
if x > y {
    // Instruction here
}
else if !(z < y) {
    // Instruction here
}
else {
    // Instruction here
}
```

Além disso, é possível inicializar algo dentro do if:
```
if x := 0; x < y {
    // Instruction here
}
```

É possível utilizar switch também da forma abaixo:
```
switch x {
    case 0, 1: 
        // Instruction
    case 2, (2 < x): 
        // Instruction
    default:
        // Instruction
}
```
A separação com vírgula indica uma possibilidade a mais.