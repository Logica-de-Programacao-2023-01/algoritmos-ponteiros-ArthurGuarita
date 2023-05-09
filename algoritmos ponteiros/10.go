package main

import "fmt"

//Implemente uma função que receba um ponteiro para uma slice e seu tamanho e preencha-o com os n primeiros números primos.

func main() {
	var s []int
	n := 5
	var p *[]int = &s
	primesN(p, n)
	fmt.Print(*p)
}

func primesN(p *[]int, n int) {
	Isprime := func(n int) bool {
		if n <= 1 {
			return false
		}
		for i := 2; i <= n; i++ {
			if n%i == 0 && i != n {
				return false
			}
		}
		return true
	}
	count := 0
	num := 2
	for {
		if Isprime(num) {
			*p = append(*p, num)
			count++
		}
		if count >= n {
			break
		}
		num++
	}
}
