package main

import "fmt"

type Conta struct {
	saldo   float64
	titular string
}

func main() {
	c := Conta{
		saldo:   0,
		titular: "Bobby",
	}
	var p *Conta = &c
	addSaldo(p)
	fmt.Print(c.titular, "\n", c.saldo)
}

func addSaldo(p *Conta) {
	newSaldo := 670.99
	(*p).saldo = newSaldo
}
