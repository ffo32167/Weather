package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	c "github.com/ffo32167/weather/weatherLogin/config"
	r "github.com/ffo32167/weather/weatherLogin/routes"
	t "github.com/ffo32167/weather/weatherLogin/templates"

	// используем логер из weatherParser
	l "github.com/ffo32167/weather/weatherParser/logger"

	"github.com/sirupsen/logrus"
)

// 1. тесты
// 2. в readme есть плашки про тестирование и доку
// 3. ci настроен правильно, и показывает покрытие
// 4. файл конфигурации CI без отсебятины
// 5. CI вызывает мета-линтер golangci-lint

// сделать возможность смены grpc сервера на лету
// добавить проверку работоспособности grpc сервера перед выполнением запроса
func main() {
	// Инициализировать конфиг
	cfg := c.NewConfig()
	// Инициализировать логер
	l.NewLog("weatherLogin", cfg.AppPath, cfg.SourceLinesInLog, cfg.LogLevel)
	// Инициализировать шаблоны
	t.Initialize(cfg.AppPath)
	// Инициализировать роутер
	r := r.NewRouter(cfg.Conn)
	// Запустить сервер
	http.Handle("/", r)
	// парочка таймаутов на всякий случай
	server := &http.Server{
		Addr:         cfg.HTTPPort,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}
	// пытаемся запустить сервер
	logrus.Info("starting server at ", cfg.HTTPPort)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			logrus.Fatal("can't start server: ", err)
		}
	}()

	// слушаем сигнал на прекращение работы
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// когда сигнал приходит, выключаем за собой свет
	<-done
	logrus.Info("server stopped")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		logrus.Info("closing grpc connection")
		cfg.Conn.Close()
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		logrus.Fatal("shutdown failed:", err)
	}
	logrus.Print("server shutdown")
}
