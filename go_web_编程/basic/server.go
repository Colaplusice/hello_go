package main

import (
	"net/http"
	"fmt"
)

func sayhello_again(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("value:", v)
	}
	fmt.Fprintf(w, "hello_world_again")
}
