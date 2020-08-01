package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	// парочка таймаутов на всякий случай
	server := &http.Server{
		Addr:         cfg.HTTPPort,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}
	// слушаем сигнал на прекращение работы
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	logrus.Info("starting server at", cfg.HTTPPort)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			logrus.Fatal("can't start server:", err)
		}
	}()

	// когда сигнал приходит, выключаем за собой свет
	<-done
	logrus.Info("server stopped")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		logrus.Info("closing database connection")
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		logrus.Fatal("shutdown failed:", err)
	}
	logrus.Print("server shutdown")
}
