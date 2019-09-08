package main

import "fmt"

type gator int

func main() {
	var g1 gator = 3
	var x = 1

	x = int(g1)
	fmt.Println(x)
	fmt.Printf("Type: %T\n", x)
	fmt.Printf("Type: %T\n", g1)

	g1.greeting()
}

func (g gator) greeting() {
	fmt.Println("Hello, I am a gator")
}
