package main

import "fmt"

func main() {
	n := 7
	var ptr = &n
	fmt.Print(sumNumerosNaturais(ptr))

}

func sumNumerosNaturais(ptr *int) int {
	sum := 0
	for i := 1; i <= *ptr; i++ {
		sum += i
	}
	return sum
}