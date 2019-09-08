package main

import (
	"fmt"
	"strings"
)

type person struct {
	fName   string
	lName   string
	favFood []string
}

func main() {
	p1 := person{fName: "john", lName: "doe", favFood: []string{"chicken", "pasta"}}
	fmt.Println(p1.walk())
}

func (p person) walk() string {
	return strings.Join([]string{p.fName, "is walking"}, " ")
}
