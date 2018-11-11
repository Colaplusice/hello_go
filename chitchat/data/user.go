package data

import (
	"time"
	"fmt"
	"crypto/sha1"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	password  string
	CreatedAt time.Time
}
type Session struct {
	//Id int
}

func UserByEmail(email string) (user User, err error) {
	user = User{}

	err = Db.QueryRow("select id ,uuid, name,email,password,created_at from user where uuid=$1", email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.password, &user.CreatedAt)
	return
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
