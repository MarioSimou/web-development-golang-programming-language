package main

import (
	"fmt"
	"net"
)

func main() {
	conn, e := net.Dial("tcp", "localhost:8080")
	check(
		e)
	defer conn.Close()
	fmt.Fprintln(conn, "I dialed you")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
