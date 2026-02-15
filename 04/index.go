package main

import (
	"fmt"
)

type car struct {
	id         int
	marca      string
	modelo     string
	cor        string
	estaLigado bool
}

func soma(x ...int) (int, int) {
	soma := 0

	for _, value := range x {
		soma += value
	}

	return soma, len(x)
}

func main() {
	carro := car{
		id:         0,
		marca:      "Marca não sei",
		modelo:     "Modelo não sei",
		cor:        "Amarelo rosado",
		estaLigado: false,
	}

	fmt.Println(carro)
	fmt.Println(carro.marca)

	soma, tamanho := soma(1, 2, 3)

	// Retornará 6, 3
	fmt.Println(soma, tamanho)
}
