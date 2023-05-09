package main

import "fmt"

//Crie uma função que receba um ponteiro para uma variável como argumento e modifique o valor da variável dentro da
//função. Certifique-se de que o ponteiro aponte para uma área de memória válida e que a memória seja liberada após o uso.

func main() {
	n := 12
	var p *int = &n
	alterarValorVariavel(p)
	fmt.Print(*p)
}

func alterarValorVariavel(p *int) {
	if p != nil {
		newValue := 45
		*p = newValue
	} else {
		fmt.Println("Ponteiro é nil")
	}
}
