package main

import (
	"fmt"
	"net/http"
)

func authReq(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("authReq call")
		// Проверить наличие сессии
		session, _ := store.Get(r, "session")
		_, ok := session.Values["user"]
		fmt.Println("ok:", ok)
		if !ok {
			fmt.Println("checkSession err")
			http.Redirect(w, r, "/login", 302)
			return
		}
		handler.ServeHTTP(w, r)
	}
}
