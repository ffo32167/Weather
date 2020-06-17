package main

import (
	"net/http"
)

// Обработчик выхода
func pageLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	delete(session.Values, "user")
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusFound)
}
