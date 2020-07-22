package main

import (
	"net/http"
	"time"

	c "github.com/ffo32167/weather/weatherLogin/config"
	h "github.com/ffo32167/weather/weatherLogin/routes"
	t "github.com/ffo32167/weather/weatherLogin/templates"

	// используем парсеровский логер
	l "github.com/ffo32167/weather/weatherParser/logger"

	"github.com/sirupsen/logrus"
)

func main() {
	// Инициализировать конфиг
	cfg := c.NewConfig()
	// Инициализировать логер
	l.NewLog("weatherLogin", cfg.AppPath, cfg.SourceLinesInLog, cfg.LogLevel)
	// Инициализировать шаблоны
	t.Initialize(cfg.AppPath)
	// Инициализировать роутер
	r := h.NewRouter()
	// Запустить сервер
	http.Handle("/", r)

	server := &http.Server{
		Addr:         cfg.HTTPPort,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}
	logrus.Info("starting server at", cfg.HTTPPort)

	if err := server.ListenAndServe(); err != nil {
		logrus.Fatal("can't start server:", err)
	}
}
