package logger

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

// NewLog настраивает логгер
func NewLog(appName string, appPath string, sourceLinesInLog *bool, logLevel string) {
	logrus.StandardLogger().Out = os.Stdout
	file, err := os.OpenFile(filepath.Join(appPath, appName+`.log`), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logrus.StandardLogger().Out = file
	} else {
		logrus.Info("Failed to write to the file, using default stderr")
	}
	logrus.StandardLogger().SetLevel(logrus.DebugLevel)
	FormatLog(*sourceLinesInLog)
	SetLogLevel(logLevel)
}

// SetLogLevel Установить уровень логирования
func SetLogLevel(LogLevel string) {
	level, err := logrus.ParseLevel(LogLevel)
	if err != nil {
		logrus.Error("Failed to parse log level, using DebugLevel")
		logrus.StandardLogger().SetLevel(logrus.DebugLevel)
	}
	logrus.StandardLogger().SetLevel(level)
}

// FormatLog Задать формат лога
func FormatLog(SourceLinesInLog bool) {
	logrus.SetReportCaller(true)
	logrus.StandardLogger().Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			s := strings.Split(f.Function, ".")
			funcname := s[len(s)-1]
			_, filename := path.Split(f.File)
			return funcname, filename + ":" + strconv.Itoa(f.Line)
		},
	}
}
