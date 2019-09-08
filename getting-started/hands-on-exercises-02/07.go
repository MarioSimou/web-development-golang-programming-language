package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle   vehicle
	fourWheel bool
}

type sedan struct {
	vehicle vehicle
	luxury  bool
}

type transportation interface {
	transportationDevice() string
}

func main() {
	t1 := truck{
		vehicle:   vehicle{doors: 2, color: "black"},
		fourWheel: true,
	}

	s1 := sedan{
		vehicle: vehicle{doors: 4, color: "blue"},
		luxury:  true,
	}

	fmt.Println(report(t1))
	fmt.Println(report(s1))
}

func (t truck) transportationDevice() string {
	return "I transfer stuff"
}

func (s sedan) transportationDevice() string {
	return "I transfer people"
}

func report(tr transportation) string {
	return tr.transportationDevice()
}
