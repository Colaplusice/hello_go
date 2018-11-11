package main

import "net/http"
import "chitchat/data"

//验证
func authenticate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, _ := data.UserByEmail(r.PostFormValue("email"))

	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		session := user.CreateSession()
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		//设置response header的cookie
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "login", 302)
	}
}
func login(writer http.ResponseWriter, request, r *http.Request) {
	t := parseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(writer, nil)
}
func logout(writer http.ResponseWriter, request, r *http.Request) {

}

func signup(writer http.ResponseWriter, request, r *http.Request) {

}
func signupAccount(writer http.ResponseWriter, request, r *http.Request) {

}
