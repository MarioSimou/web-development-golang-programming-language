package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	ln, e := net.Listen("tcp", ":8080")
	check(e)
	defer ln.Close()

	for {
		conn, e := ln.Accept()
		check(e)
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		t := scanner.Text()
		fmt.Fprintln(conn, "Server says: "+t)
	}

	fmt.Println("finish scanning...")
	conn.Close()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
