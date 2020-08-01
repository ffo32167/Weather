package middleware

import (
	"net/http"

	s "github.com/ffo32167/weather/weatherLogin/session"
	"github.com/sirupsen/logrus"
)

// AuthReq это миддлваре, проверяющее аутентифицирован ли пользователь
func AuthReq(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверить наличие сессии
		session, err := s.Store.Get(r, "session")
		if err != nil {
			logrus.Info("can't decode session")
		}
		_, ok := session.Values["user"]
		if !ok {
			http.Redirect(w, r, "/login", 302)
			return
		}
		handler.ServeHTTP(w, r)
	}
}

// AuthReqHandler это миддлваре, проверяющее аутентифицирован ли пользователь
func AuthReqHandler(handler http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// Проверить наличие сессии
		session, err := s.Store.Get(r, "session")
		if err != nil {
			logrus.Info("can't decode session")
		}
		_, ok := session.Values["user"]
		if !ok {
			http.Redirect(w, r, "/login", 302)
			return
		}
		handler.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
