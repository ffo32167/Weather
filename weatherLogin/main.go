package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	config := newConfig()
	sessManager = newSessionManager()
	r := mux.NewRouter()
	r.HandleFunc("/", config.pageWeatherGet).Methods("GET")
	r.HandleFunc("/", config.pageWeatherPost).Methods("POST")
	r.HandleFunc("/login", pageLogin)
	r.HandleFunc("/logout", pageLogout)
	http.Handle("/", r)

	server := &http.Server{
		Addr:         config.HTTPPort,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}
	fmt.Println("starting server at", config.HTTPPort)

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("couldn't start server:", err)
		log.Fatal(err)
	}
}
