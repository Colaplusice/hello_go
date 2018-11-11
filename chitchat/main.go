package main

import (
	"net/http"
	"text/template"
	"fmt"
)

func main() {

	mux := http.NewServeMux()
	// 构造一个 handler 来管理文件路由
	files := http.FileServer(http.Dir("/public"))

	//路由
	mux.Handle("/static", http.StripPrefix("/static", files))
	mux.HandleFunc("/", index_1)
	//mux.HandleFunc("/error", err)

	////auth route
	//mux.HandleFunc("/login", login)
	//mux.HandleFunc("/logout", logout)
	//mux.HandleFunc("/signup", signup)
	//mux.HandleFunc("/signup_account", signupAccount)
	//mux.HandleFunc("/authenticate", authenticate)
	//
	//// 发帖相关的
	//mux.HandleFunc("/thread/new", newThread)
	//mux.HandleFunc("/thread/create", createThread)
	//mux.HandleFunc("/thread/post", postThread)
	//mux.HandleFunc("/thread/read", readThread)

	//构造一个server对象
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

func generate_html(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	//将所有的filename组合到一起
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)

}
