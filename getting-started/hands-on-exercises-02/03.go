package main

import "fmt"

type person struct {
	fName string
	lName string
}

func main() {
	p1 := person{fName: "john", lName: "doe"}
	fmt.Println(p1.fName)
	fmt.Println(p1.lName)
}
