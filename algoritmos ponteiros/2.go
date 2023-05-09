package main

import "fmt"

func main() {
	n := 23
	var p = &n
	fmt.Print(IsParORImpar(p))
}

func IsParORImpar(p *int) string {
	if *p%2 == 0 {
		return "É par"
	} else {
		return "Ímpar"
	}
}
