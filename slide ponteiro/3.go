package main

import "fmt"

type Produto struct {
	nome       string
	preco      float64
	quantidade int
}

func main() {
	c := Produto{
		nome:       "Samsung A23",
		preco:      1400.99,
		quantidade: 19,
	}
	var ptr = &c.preco
	changePrice(ptr)
	fmt.Print(*ptr)
}

func changePrice(ptr *float64) {
	newPRice := 2100.99
	*ptr = newPRice
}
