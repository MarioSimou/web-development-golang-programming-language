package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

type human interface {
	greetings()
}

func saySomething(h human) {
	fmt.Println("What should i say?")
}

func main() {
	p1 := person{fname: "john", lname: "doe"}
	sa1 := secretAgent{
		person:        person{fname: "James", lname: "Bond"},
		licenseToKill: true,
	}

	p1.greetings()
	sa1.greetings()
	sa1.person.greetings()
	saySomething(sa1)
}

func (sa secretAgent) greetings() {
	fmt.Printf("Hello, my name is %v %v. Shaken, not stirred\n", sa.fname, sa.lname)
}

func (p person) greetings() {
	fmt.Printf("Hello, my name is %v %v\n", p.fname, p.lname)
}
