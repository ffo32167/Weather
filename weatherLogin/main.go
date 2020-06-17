package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
)

var templates *template.Template

func main() {
	config := newConfig()
	templatesPath := filepath.Join(config.appPath, `templates`, `*.html`)
	t, err := template.ParseGlob(templatesPath)
	if err != nil {
		fmt.Println("template.ParseGlob err:", err)
	}
	templates = template.Must(t, err)

	r := mux.NewRouter()
	r.HandleFunc("/", authReq(pageWeatherGet)).Methods("GET")
	r.HandleFunc("/", authReq(pageWeatherPost)).Methods("POST")
	r.HandleFunc("/login", pageLoginGet).Methods("GET")
	r.HandleFunc("/login", pageLoginPost).Methods("POST")
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
