package main

import (
	ch "github.com/ffo32167/weather/weatherParser/cache"
	c "github.com/ffo32167/weather/weatherParser/config"
	g "github.com/ffo32167/weather/weatherParser/grpcservice"
	l "github.com/ffo32167/weather/weatherParser/logger"
	"github.com/sirupsen/logrus"
)

// todo
// добавить кэш работающий на Redis

func main() {
	// прочитать конфиг
	cfg, err := c.NewConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	// настроить логер
	l.NewLog("weatherParser", cfg.AppPath, cfg.SourceLinesInLog, cfg.LogLevel)

	// создать кэш
	logrus.Info("Reading Cache")
	wmc := ch.NewWeatherCache(cfg.AppPath)

	// запустить сервер
	logrus.Info("GRPC service starting up...")
	g.ServerStart(cfg, wmc)
}
