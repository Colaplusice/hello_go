package main

import (
	"net/http"
	"chitchat/data"
	"fmt"
)

//生成html 写入reponsewriter
func index(writer http.ResponseWriter, request *http.Request) {
	threads, err := data.Threads()
	var templates *template.Template
	if err != nil {
		error_message(writer, request, "can't not open thread")
		//templates = template.Must(template.ParseFiles(private_tmp_files...))
	} else {
		_, err := session(writer, request)
		if err != nil {
			generate_html(writer, request, "layout", "public.navbar", "index")
		} else {
			generate_html(writer, request, "layout", "private.navbar", "index")

		}
	}
	templates.ExecuteTemplate(w, "layout", threads)
	thread, err := data.Threads();
	if err == nil {
		//解析layout  html
		templates.ExecuteTemplate(w, "layout", thread)

	}

}

func index_1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
