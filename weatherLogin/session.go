package main

import (
	"github.com/gorilla/sessions"
)

type user struct {
	username string
}

var store = sessions.NewCookieStore([]byte("cookie-store"))

func authUser(login, password string) (*user, error) {
	user, err := getUserByLogin(login)
	return user, err
}

// впускаем всех
func getUserByLogin(login string) (*user, error) {
	return &user{
		username: "user_" + login,
	}, nil
}
