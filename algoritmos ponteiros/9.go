package main

import "fmt"

//Implemente uma função que receba um ponteiro para uma struct "Livro" com campos "Título", "Autor" e "Preço", e que
//adicione um desconto de 10% no preço do livro.

type Livro1 struct {
	titulo string
	autor  string
	preco  float64
}

func main() {
	c := Livro1{
		titulo: "Ordem Paranormal RPG",
		autor:  "Cellbit",
		preco:  129.90,
	}
	var p = &c
	descount10porCent(p)
	fmt.Printf("Com o desconto, o preço é %.2f", (*p).preco)
}

func descount10porCent(p *Livro1) {
	(*p).preco = (*p).preco * 0.9
}
