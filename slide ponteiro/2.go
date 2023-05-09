package main

import "fmt"

func main() {
	t := true
	var ptr = &t
	invertBool(ptr)
	fmt.Print(*ptr)
}

func invertBool(ptr *bool) {
	*ptr = !*ptr
}
