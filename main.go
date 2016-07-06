package main

import (
	"fmt"
	"net/http"
	"golang.org/x/net/context"

	"goji.io"
	"goji.io/pat"
)

func hello(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	name := pat.Param(ctx, "name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	mux := goji.NewMux()
	mux.HandleFuncC(pat.Get("/hello/:name"), hello)

	http.ListenAndServe("localhost:8000", mux)
}



