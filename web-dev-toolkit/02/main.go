package main

import (
	"encoding/base64"
	"os"
)

func main() {
	w := base64.NewEncoder(base64.StdEncoding, os.Stdout) // creates a writer that is returned
	w.Write([]byte("sometest?;"))
	defer w.Close()
}
