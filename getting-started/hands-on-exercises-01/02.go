package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person person
	secret bool
}

func main() {
	p1 := person{fname: "John", lname: "Doe"}
	sa := secretAgent{
		person: person{fname: "James", lname: "Bond"},
		secret: true,
	}

	fmt.Println(p1.fname)
	p1.pSpeak()

	fmt.Println(sa.person.fname)
	sa.saSpeak()
}

func (p person) pSpeak() {
	fmt.Printf("Hello, my name is %v %v.\n", p.fname, p.lname)
}

func (sa secretAgent) saSpeak() {
	fmt.Printf("Hello, my name is %v %v. I am a secret agent.\n", sa.person.fname, sa.person.lname)
}
