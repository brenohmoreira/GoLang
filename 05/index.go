package main

import (
	"fmt"
)

func main() {
	x := 10

	// Isso retorna endereço de X
	fmt.Println(&x)

	// Salva endereço de x em y
	y := &x

	// Busca valor existente em um endereço
	fmt.Println(*y)

	// Adicionando 1 no valor de do endereço de x
	*y++

	// Não adicionei diretamente em x, adicionei por meio do ponteiro y
	fmt.Println(x)

	// *y == x
}
