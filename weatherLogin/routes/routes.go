package routes

import (
	m "github.com/ffo32167/weather/weatherLogin/middleware"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

// NewRouter инициализирует роутер
func NewRouter(conn *grpc.ClientConn) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", m.AuthReq(pageWeatherGet)).Methods("GET")
	r.Handle("/", m.AuthReqHandler(pageWeatherHandler{conn})).Methods("POST")
	r.HandleFunc("/login", pageLoginGet).Methods("GET")
	r.HandleFunc("/login", pageLoginPost).Methods("POST")
	r.HandleFunc("/logout", pageLogout)
	return r
}
