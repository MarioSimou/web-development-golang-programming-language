package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	tcp, e := net.Listen("tcp", ":8080")
	check(e)

	for {
		conn, e := tcp.Accept()
		check(e)
		go handle(conn)
	}

}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := strings.ToLower(scanner.Text())
		fmt.Fprintf(conn, "%s <-> %s\n", ln, strings.Join(rot13(ln), ""))
	}
}

func rot13(s string) []string {
	bytes := []byte(s)
	sl := make([]string, len(bytes))
	for i, b := range bytes {
		// 97 - 122 (a - z)
		nb := b + 13
		if nb > 122 {
			nb = nb - 26
		}
		sl[i] = string(nb)
	}

	return sl
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
