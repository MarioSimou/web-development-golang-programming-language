package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, "userId", 12)
	result, e := dbAccess(ctx)
	if e != nil {
		http.Error(w, e.Error(), http.StatusRequestTimeout)
	}

	fmt.Fprintln(w, result)

}

func dbAccess(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	ch := make(chan int)
	go func() {
		uuid := ctx.Value("userId").(int)
		time.Sleep(10 * time.Second)

		if ctx.Err() != nil {
			return
		}

		ch <- uuid
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case i := <-ch:
		return i, nil
	}
}

func main() {
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
