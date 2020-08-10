package routes

import (
	"net/http"

	s "github.com/ffo32167/weather/weatherLogin/session"
	"github.com/sirupsen/logrus"
)

// Обработчик выхода
func pageLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := s.Store.Get(r, "session")
	delete(session.Values, "user")
	err := session.Save(r, w)
	if err != nil {
		logrus.Error("can't save session", err)
	}
	http.Redirect(w, r, "/login", http.StatusFound)
}
