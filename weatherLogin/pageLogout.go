package main

import (
	"net/http"
	"time"
)

// Обработчик выхода
func pageLogout(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	// Если уже нет cookie
	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/", http.StatusFound)
		return
		// Если другая ошибка
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Убрать cookie
	sessManager.Delete(&sessionID{
		ID: session.Value,
	})
	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
	// Указать на выход
	http.Redirect(w, r, "/", http.StatusFound)
}
