package main

import (
	"fmt"
	"errors"
	)
	
func main() {

	pilha := Pilha{}
	fmt.Printf("Pilha criada com tamanho %d\n", pilha.Tamanho())
	fmt.Printf("Vazia? %v\n\n", pilha.Vazia())
	
	pilha.Empilhar("Go")
	pilha.Empilhar(2009)
	pilha.Empilhar(3.14)
	pilha.Empilhar("Fim")
	fmt.Printf("Tamanho após empilhar 4 valores: %d\n", pilha.Tamanho())
	fmt.Printf("Vazia? %v\n\n", pilha.Vazia())
	
	for !pilha.Vazia() {
		v, _ := pilha.Desempilhar()
		fmt.Printf("Desempilhando %v\n", v)
		fmt.Printf("Tamanho: %d\n", pilha.Tamanho())
		fmt.Printf("Vazia? %v\n\n", pilha.Vazia())
	}
	
	_, err := pilha.Desempilhar()
	if err != nil {
		fmt.Printf("Erro: %s\n", err)
	}

}

type Pilha struct {
	valores []interface{}
}

func (pilha Pilha) Tamanho() int {
	return len(pilha.valores)
}

func (pilha Pilha) Vazia() bool {
	return pilha.Tamanho() == 0
}

func (pilha *Pilha) Empilhar(valor interface{}) {
	pilha.valores = append(pilha.valores, valor)
}

func (pilha *Pilha) Desempilhar() (interface{}, error) {
	if pilha.Vazia() {
		return nil, errors.New("Pilha Vazia !!!")
	}
	valor := pilha.valores[pilha.Tamanho()-1]
	pilha.valores = pilha.valores[:pilha.Tamanho()-1]
	return valor, nil
}