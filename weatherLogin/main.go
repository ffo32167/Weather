package main

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	config := newConfig()
	sessManager = newSessionManager()
	http.HandleFunc("/", config.pageInner)
	http.HandleFunc("/login", pageLogin)
	http.HandleFunc("/logout", pageLogout)
	fmt.Println("starting server at", config.HTTPPort)
	log.WithFields(logrus.Fields{"HTTP server at port": config.HTTPPort}).Error("starting HTTP server")
	http.ListenAndServe(config.HTTPPort, nil)
}
