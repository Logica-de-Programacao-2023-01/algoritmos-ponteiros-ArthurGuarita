package main

import "fmt"

type Livro struct {
	titulo string
	autor  string
}

func main() {
	c := Livro{
		titulo: "Harry Potter",
		autor:  "Anônimo",
	}
	var p *Livro = &c
	alterarTitulo(p)
	fmt.Print(c.titulo, "\n", c.autor)
}

func alterarTitulo(p *Livro) {
	if (*p).autor == "Anônimo" {
		(*p).titulo = "Desconhecido"
	}
}
