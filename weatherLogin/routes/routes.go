package routes

import (
	m "github.com/ffo32167/weather/weatherLogin/middleware"
	"github.com/gorilla/mux"
)

// NewRouter инициализирует роутер
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", m.AuthReq(PageWeatherGet)).Methods("GET")
	r.HandleFunc("/", m.AuthReq(PageWeatherPost)).Methods("POST")
	r.HandleFunc("/login", PageLoginGet).Methods("GET")
	r.HandleFunc("/login", PageLoginPost).Methods("POST")
	r.HandleFunc("/logout", PageLogout)
	return r
}
