package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 1234
	var p *int = &n
	fmt.Print(sumTwoDigits(p))

}

func sumTwoDigits(p *int) int {
	c := strconv.Itoa(*p)
	s1 := c[len(c)-1]
	s2 := c[len(c)-2]
	m, err := strconv.Atoi(string(s1))
	if err != nil {
		fmt.Errorf("ocorreu um erro %s", err)
	}
	n, er := strconv.Atoi(string(s2))
	if er != nil {
		fmt.Errorf("ocorreu um erro %s", er)
	}
	sum := n + m
	return int(sum)
}
