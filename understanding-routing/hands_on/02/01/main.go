package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for true {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	var method string
	var path string

	for scanner.Scan() {
		t := scanner.Text()
		// fmt.Println("I see your string... ", t)

		if strings.Contains(t, "HTTP/1.1") {
			slice := strings.Split(t, " ")
			method = slice[0]
			path = slice[1]
		}
		if t == "" {
			// when ln is empty, header is done
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break
		}
	}

	switch {
	case path == "/" && method == "GET":
		renderHTML("this is the home page", conn)
	case path == "/apply" && method == "GET":
		renderHTML(renderForm("/apply"), conn)
	case path == "/apply" && method == "POST":
		renderHTML("handling.. post", conn)
	default:
		renderHTML("Welcome", conn)
	}

	defer conn.Close()
}

func renderForm(action string) string {
	return `
	<form method="POST" action="` + action + `" enctype="application/x-www-form-urlencoded">
		<div>
			<label for="name">Name:</label>
			<input type="text" name="name" id="name" placeholder="Your name" />
		</div>
		<div>
			<label for="age">Age:</label>
			<input type="text" name="age" id="age" placeholder="Your age" />
		</div>
		<div>
			<button>Submit</button>
		</div>
	</form>`
}

func renderHTML(c string, conn net.Conn) {
	doc := "<!DOCTYPE html><html><head><title>Home</title></head><body>$</body></html>"
	upDoc := strings.Replace(doc, "$", c, -1)

	fmt.Fprintln(conn, "HTTP/1.1 200 OK")
	fmt.Fprintln(conn, "Content-Type: text/html")
	fmt.Fprintln(conn, "Content-Encoding: utf-8")
	fmt.Fprintln(conn, "Content-Length:", len(upDoc))
	fmt.Fprintln(conn, "")
	fmt.Fprintln(conn, upDoc)
}
