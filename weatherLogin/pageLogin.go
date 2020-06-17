package main

import (
	"fmt"
	"net/http"
)

func pageLoginGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("loginGet call")
	// проверяем есть ли активная сессия у клиента
	session, _ := store.Get(r, "session")
	_, ok := session.Values["user"]
	if ok {
		http.Redirect(w, r, "/", 302)
		return
	}
	// парсим шаблон
	if err := templates.ExecuteTemplate(w, "login.html", nil); err != nil {
		log.Error("can't parse login.html:", err)
	}
}

// Обработчик страницы логина
func pageLoginPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("pageLoginPost call")
	// Значение поля Логин из html-формы
	r.ParseForm()
	inputLogin := r.PostForm.Get("login")
	inputPassword := r.PostForm.Get("password")
	fmt.Println("pageLoginPost login:", inputLogin)
	fmt.Println("pageLoginPost password:", inputPassword)
	user, err := authUser(inputLogin, inputPassword)
	if err != nil {
		fmt.Println("authUser error:", err)
	}

	session, err := store.Get(r, "session")
	// Если что-то пошло не так
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	session.Values["user"] = user.username
	session.Save(r, w)
	// и отправить на главную
	http.Redirect(w, r, "/", http.StatusFound)
}
