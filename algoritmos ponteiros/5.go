package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64 = 67.34
	var p *float64 = &n
	mediaPont(p)
	fmt.Print(*p)
}
func mediaPont(p *float64) {
	*p = *p + math.Pi/2
}
