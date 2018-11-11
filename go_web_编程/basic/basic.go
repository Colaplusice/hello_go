package main

import (
	"net/http"
	"log"
)


func main() {
	//注册路由
	http.HandleFunc("/hello_again", sayhello_again)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("listen_server", err)
	}
}
