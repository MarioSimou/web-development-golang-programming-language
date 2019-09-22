package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	tcp, e := net.Listen("tcp", ":8080")
	check(e)
	defer tcp.Close()

	for {
		conn, e := tcp.Accept()
		check(e)
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println("Line:", ln)
		fmt.Fprintln(conn, "I got it! I have your request")
	}

	defer conn.Close()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
