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

		go handleRequest(conn)
	}
}
func handleRequest(conn net.Conn) {
	fmt.Println(conn.LocalAddr().String())

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		t := scanner.Text()
		fmt.Println("Text: ", t)
	}

	fmt.Println("finish scanning...")

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
