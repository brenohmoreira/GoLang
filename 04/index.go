package main

import (
	"fmt"
)

type car struct {
	id  int
	cor string
}

func (c car) printCar(name string) {
	fmt.Printf("ID: %d\nNOME: %s\nCOR: %s\n", c.id, c.cor, name)
}

func soma(x ...int) (int, int) {
	soma := 0

	for _, value := range x {
		soma += value
	}

	return soma, len(x)
}

// Interfaces
type triangulo struct {
	medidaLado float64
	medidaBase float64
}

type quadrado struct {
	medidaLado float64
}

type retangulo struct {
	medidaLado      float64
	medidaOutroLado float64
}

func (t triangulo) calculaArea() float64 {
	return (t.medidaBase * t.medidaLado) / 2
}

func (q quadrado) calculaArea() float64 {
	return q.medidaLado * q.medidaLado
}

func (r retangulo) calculaArea() float64 {
	return r.medidaOutroLado * r.medidaLado
}

// Todo tipo (objeto/classe) que implicitamente implementa calculaArea(), implementa figura
type figura interface {
	calculaArea() float64
}

// Quadrado, triangulo e retangulo implicitamente implementam figura e são figuras, pois elas cumprem o contrato implementando as suas próprias calculaArea()
// Ao chamar esse método enviando uma figura (interface) e enviando o que implementa ele, temos uma mesma função chamando funções com a regra de negócio certa a depender do tipo enviado
func mostraArea(f figura) float64 {
	return f.calculaArea()
}

func main() {
	carro := car{id: 0, cor: "Branco"}

	carro.printCar("Mercedes")

	soma, tamanho := soma(1, 2, 3)

	// Retornará 6, 3
	fmt.Println(soma, tamanho)

	// Interfaces
	figuraQuadrado := quadrado{medidaLado: 5}
	figuraRetangulo := retangulo{medidaLado: 5, medidaOutroLado: 10}
	figuraTriangulo := triangulo{medidaLado: 5, medidaBase: 7}

	areaQuadrado := mostraArea(figuraQuadrado)
	areaRetangulo := mostraArea(figuraRetangulo)
	areaTriangulo := mostraArea(figuraTriangulo)

	fmt.Println(areaQuadrado, areaRetangulo, areaTriangulo)
}
