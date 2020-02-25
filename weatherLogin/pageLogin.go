package main

import (
	"net/http"
	"time"
)

// Обработчик страницы логина
func pageLogin(w http.ResponseWriter, r *http.Request) {
	// Значение поля Логин из html-формы
	inputLogin := r.FormValue("login")
	// Срок годности cookie
	expiration := time.Now().Add(time.Hour)
	// Создаём и заполняем структуру Сессия
	sess, err := sessManager.Create(&session{
		Login: inputLogin,
	})
	// Если что-то пошло не так
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Создать и заполнить cookie
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   sess.ID,
		Expires: expiration,
	}
	// Проставить штампик
	http.SetCookie(w, &cookie)
	// и отправить на главную
	http.Redirect(w, r, "/", http.StatusFound)
}
