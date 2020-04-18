package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	config := newConfig()
	sessManager = newSessionManager()
	r := mux.NewRouter()
	r.HandleFunc("/", config.pageInnerGet).Methods("GET")
	r.HandleFunc("/", config.pageInnerPost).Methods("POST")
	http.Handle("/", r)

	http.HandleFunc("/login", pageLogin)
	http.HandleFunc("/logout", pageLogout)
	fmt.Println("starting server at", config.HTTPPort)
	log.WithFields(logrus.Fields{"HTTP server at port": config.HTTPPort}).Error("starting HTTP server")
	log.Fatal(http.ListenAndServe(config.HTTPPort, nil))
}
