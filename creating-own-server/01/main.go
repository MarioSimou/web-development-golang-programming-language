package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	li, e := net.Listen("tcp", ":8080")
	check(e)

	defer li.Close()

	for {
		conn, e := li.Accept()
		check(e)

		fmt.Println(conn.LocalAddr().String())
		fmt.Println(conn.RemoteAddr().String())

		io.WriteString(conn, "HELLO WOLRD")
		fmt.Fprintln(conn, "Hello world twice")

		conn.Close()
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
