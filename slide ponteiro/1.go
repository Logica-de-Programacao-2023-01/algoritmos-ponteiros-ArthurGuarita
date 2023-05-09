package main

import "fmt"

func main() {
	a := 12
	b := 34
	var ptr1 = &a
	var ptr2 = &b
	swap(ptr1, ptr2)
	fmt.Print(*ptr1, *ptr2)
}

func swap(ptr1 *int, ptr2 *int) {
	var aux = *ptr1
	*ptr1 = *ptr2
	*ptr2 = aux
}
