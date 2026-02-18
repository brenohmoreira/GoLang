## Ponteiro

Um ponteiro é um tipo que apontará para o endereço de uma variável.

- & busca endereço de algo
- \* busca valor em um endereço

Quando passamos para uma função um valor x, temos alguns problemas: ao executar a função, é criado uma cópia de x <br>
Essa cópia gera dois problemas:
- Performance, pois há a criação de uma variável de forma desnecessária
- Mutabilidade, pois dentro da função a cópia não vai alterar o valor da variável original

Se não for esse o objetivo, se faz:
```
func main() {
    x += 10
    x = adicionaQuantidade(&x, 10)
    
    // Será 20
    fmt.Printf(x)     
}

// x é o ponteiro, pois aponta para a variável original (*int)
func adicionaQuantidade(x *int, value x) int {
    // Valor da variável apontada por x + value
    *x += value
    return *x
}
```
