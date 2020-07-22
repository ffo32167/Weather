package routes

import (
	"net/http"

	s "github.com/ffo32167/weather/weatherLogin/session"
)

// PageLogout Обработчик выхода
func PageLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := s.Store.Get(r, "session")
	delete(session.Values, "user")
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusFound)
}
