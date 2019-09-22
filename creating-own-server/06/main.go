package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type header struct {
	name  string
	value string
}

type mapping struct {
	name  string
	value string
}
type mappings []mapping

func main() {
	tcp, e := net.Listen("tcp", ":8080")
	check(e)
	defer tcp.Close()

	for {
		conn, e := tcp.Accept()
		check(e)
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	response(parseRequest(conn))
	conn.Close()
}

func parseRequest(conn net.Conn) (net.Conn, mappings) {
	m := make(mappings, 0)
	c := 0
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		if c == 0 {
			arr := strings.Split(ln, " ")
			m = append(m, mapping{name: "Method", value: arr[0]})
			m = append(m, mapping{name: "Path", value: arr[1]})
		}

		if ln == "" {
			break
		}

		c++
	}

	return conn, m
}

func response(conn net.Conn, m mappings) {
	// sends response line
	fmt.Fprintln(conn, `HTTP/1.1 200 OK`)
	// follows by sendings the headers
	for _, header := range []header{
		header{name: "Content-Type", value: "text/html"},
		header{name: "Expires", value: "-1"},
		header{name: "Cache-Control", value: "no-cache"},
	} {
		fmt.Fprintln(conn, header.toString())
	}

	// sends empty line (indicates termination of the response headers)
	fmt.Fprintln(conn, "")
	// sends response body
	html := `<!DOCTYPE html><html><head><title>Tcp Server for HTTP</title></head><body><nav><a href="/home" style="display:block;">Home</a><a href="/about" style="display:block;">About</a><a href="/contact" style="display:block;">Contact</a></nav><p>$</p></body></html>`
	fmt.Fprintln(conn, strings.Replace(html, "$", m.toHTML(), 1))
}

func (h header) toString() string {
	return strings.Join([]string{h.name, h.value}, ":")
}

func (m mapping) toHTML() string {
	return "<li>" + m.name + " is " + m.value + "</li>"
}

func (m mappings) toHTML() string {
	s := make([]string, len(m))
	for i, v := range m {
		s[i] = v.toHTML()
	}

	return "<ul>" + strings.Join(s, "\n") + "</ul>"
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
