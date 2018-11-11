package main

import (
	"net/http"
	"fmt"
	"log"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Print(r.Form)
	fmt.Print(r.URL.Path)
	fmt.Print(r.URL.Host)
	fmt.Fprintf(w,"hello go")
}
func main() {
	http.HandleFunc("/", sayhello)
	err:=http.ListenAndServe(":9090",nil)
	if err!=nil{
		log.Fatal("listernAndServe:",err)
	}

}
