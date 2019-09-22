package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, e := os.Open("./table.csv")
	check(e)

	bf := make([]byte, 32*1024) // 32kilobytes
	for {
		n, e := f.Read(bf)
		if n > 0 {
			fmt.Println(string(bf[:n]))
		}

		if e == io.EOF {
			break
		}

		check(e)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
