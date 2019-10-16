package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, "uuid", "1234")
	uuid := getUuid(ctx)

	fmt.Fprintln(w, uuid)
}

func getUuid(ctx context.Context) string {
	s := ctx.Value("uuid").(string)
	fmt.Println(s)
	return s
}

func main() {
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
