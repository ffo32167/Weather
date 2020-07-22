package session

import (
	"github.com/gorilla/sessions"
)

// User это пользователь
type User struct {
	Username string
}

// Store хранилище сессий
var Store = sessions.NewCookieStore([]byte("cookie-store"))

// AuthUser проверяет пользователя
func AuthUser(login, password string) (*User, error) {
	user, err := getUserByLogin(login)
	return user, err
}

// впускаем всех
func getUserByLogin(login string) (*User, error) {
	return &User{
		Username: "user_" + login,
	}, nil
}
