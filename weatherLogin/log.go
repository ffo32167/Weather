package main

import (
	"os"
	"path/filepath"

	lrf "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// Создать и настроить логгер
func newLog(appPath string) {
	log.Out = os.Stdout
	file, err := os.OpenFile(filepath.Join(appPath, "weatherLogin.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed write to file, using default stderr")
	}
	log.SetLevel(logrus.TraceLevel)
}

func setLogFormat(SourceLinesInLog bool) {
	childFormatter := logrus.TextFormatter{}
	runtimeFormatter := &lrf.Formatter{ChildFormatter: &childFormatter}
	runtimeFormatter.File = true
	runtimeFormatter.Line = SourceLinesInLog
	log.Formatter = runtimeFormatter
}

// Установить уровень логирования
func setLogLevel(LogLevel string) {
	level, err := logrus.ParseLevel(LogLevel)
	if err != nil {
		log.Error("Failed to parse log level, using DebugLevel")
		log.SetLevel(logrus.DebugLevel)
	}
	log.SetLevel(level)
}
