package main

import (
	"fmt"
	"sort"

	"golang.org/x/crypto/bcrypt"
)

type Pessoa struct {
	Nome  string
	Idade int
}

type PorIdade []Pessoa

func (p PorIdade) Len() int           { return len(p) }
func (p PorIdade) Less(i, j int) bool { return p[i].Idade < p[j].Idade }
func (p PorIdade) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	pessoas := []Pessoa{
		{"Ana", 25},
		{"Carlos", 20},
		{"Bruno", 30},
	}

	// Ordenar por idade (crescente)
	sort.Sort(PorIdade(pessoas))

	fmt.Println(pessoas)

	// bcrypt
	password := "zanelor12"

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 4)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(hash))

	isSame := bcrypt.CompareHashAndPassword(hash, []byte("zanelsor12"))

	if isSame == nil {
		fmt.Printf("Senha correta")
	} else {
		fmt.Printf("Senha incorreta")
	}
}
