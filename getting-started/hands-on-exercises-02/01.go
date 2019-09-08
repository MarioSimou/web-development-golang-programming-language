package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3, 4}
	for i := range s1 {
		fmt.Println(i)
	}

	for i, v := range s1 {
		fmt.Println(i, v)
	}
}
