package main

import (
	"fmt"
	"time"
)

func main() {
	for n := range gen() {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}

	time.Sleep(time.Second * 1)
}

func gen() <-chan int {
	ch := make(chan int)

	go func() {
		var n int
		for {
			ch <- n
			n++
		}
	}()

	return ch
}
