package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	l, e := net.Listen("tcp", ":8080")
	check(e)
	defer l.Close()

	for {
		conn, e := l.Accept()
		check(e)

		io.WriteString(conn, "dance dance dance")
		fmt.Fprintln(conn, "shake shake")

		conn.Close()
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
