package routes

import (
	"net/http"

	s "github.com/ffo32167/weather/weatherLogin/session"
	t "github.com/ffo32167/weather/weatherLogin/templates"
	"github.com/sirupsen/logrus"
)

// PageLoginGet Обработчик Get страницы логина
func PageLoginGet(w http.ResponseWriter, r *http.Request) {
	// проверяем есть ли активная сессия у клиента
	session, err := s.Store.Get(r, "session")
	if err != nil {
		logrus.Info("can't decode session", err)
	}
	_, ok := session.Values["user"]
	if ok {
		http.Redirect(w, r, "/", 302)
		return
	}
	// парсим шаблон
	if err := t.Templates.ExecuteTemplate(w, "login.html", nil); err != nil {
		logrus.Error("can't parse login.html:", err)
	}
}

// PageLoginPost Обработчик Post страницы логина
func PageLoginPost(w http.ResponseWriter, r *http.Request) {
	// Значение поля Логин из html-формы
	r.ParseForm()
	inputLogin := r.PostForm.Get("login")
	inputPassword := r.PostForm.Get("password")
	user, err := s.AuthUser(inputLogin, inputPassword)
	if err != nil {
		logrus.Info("authUser error:", err)
	}
	session, err := s.Store.Get(r, "session")
	// Если что-то пошло не так
	if err != nil {
		logrus.Info("can't decode session", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	session.Values["user"] = user.Username
	session.Save(r, w)
	// и отправить на главную
	http.Redirect(w, r, "/", http.StatusFound)
}
