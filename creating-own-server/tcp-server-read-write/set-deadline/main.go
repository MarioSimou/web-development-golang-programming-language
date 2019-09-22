package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	m := map[string]string{
		"hello":            "hello",
		"how are you?":     "I am good. u?",
		"how old are you?": "I don't know",
	}

	l, e := net.Listen("tcp", ":8080")
	check(e)
	defer l.Close()

	for {
		conn, e := l.Accept()
		check(e)
		go handle(conn, m)
	}

}

func handle(conn net.Conn, m map[string]string) {
	e := conn.SetDeadline(time.Now().Add(time.Second * 10))
	check(e)
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		t := scanner.Text()
		fmt.Fprintln(conn, "Response: ", m[t])
	}

	defer conn.Close()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
