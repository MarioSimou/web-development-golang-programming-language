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

type human interface {
	speak()
}

func main() {
	p1 := person{fname: "John", lname: "Doe"}
	sa := secretAgent{
		person: person{fname: "James", lname: "Bond"},
		secret: true,
	}

	interfaceSpeak(p1)
	interfaceSpeak(sa)
}

func (p person) speak() {
	fmt.Printf("Hello, my name is %v %v.\n", p.fname, p.lname)
}

func (sa secretAgent) speak() {
	fmt.Printf("Hello, my name is %v %v. I am a secret agent.\n", sa.person.fname, sa.person.lname)
}

func interfaceSpeak(h human) {
	h.speak()
}
