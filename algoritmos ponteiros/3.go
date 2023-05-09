package main

import "fmt"

func main() {
	a := "Abacaxi"
	var p *string = &a
	fmt.Print(inverterStringP(p))
}

func inverterStringP(p *string) string {
	b := ""

	for i := len(*p) - 1; i >= 0; i-- {
		b += string((*p)[i])
	}
	return b
}
