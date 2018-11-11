package main

import (
	"net/http"
	"errors"
	"chitchat/data"
)

func session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error) {
	// 从request 中来找 session
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		//在数据库中找
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			// go 的异常
			err =errors.New("invalid session")
		}
	}
	return
}
