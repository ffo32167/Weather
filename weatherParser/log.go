package main

import (
	"os"
	"path/filepath"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// Настроить логер
func newLog(appPath string) {
	log.Out = os.Stdout
	file, err := os.OpenFile(filepath.Join(appPath, `weatherParser.log.json`), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to write to the file, using default stderr")
	}
	log.SetLevel(logrus.DebugLevel)
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

// Задать формат лога
func formatLog(SourceLinesInLog bool) {
	childFormatter := logrus.TextFormatter{}
	runtimeFormatter := &runtime.Formatter{ChildFormatter: &childFormatter}
	runtimeFormatter.File = true
	runtimeFormatter.Line = SourceLinesInLog
	log.Formatter = runtimeFormatter
}
