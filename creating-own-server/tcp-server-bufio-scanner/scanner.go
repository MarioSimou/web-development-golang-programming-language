package main

import (
	"bufio"
	"fmt"
	"strings"
)

var split bufio.SplitFunc

func main() {
	// s := "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
	s := "today,is,saturday,yeaaaaa"

	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		t := scanner.Text()
		fmt.Println(t)
	}

	fmt.Println("finish scanning...")
}
