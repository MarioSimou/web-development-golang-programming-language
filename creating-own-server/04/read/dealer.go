package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	d, e := net.Dial("tcp", "localhost:8080")
	check(e)

	bf := make([]byte, 80)
	for {
		_, e := d.Read(bf)
		if e == io.EOF {
			fmt.Println("finish reading...")
			break
		}
		check(e)

		fmt.Println("Read:", string(bf))

	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
