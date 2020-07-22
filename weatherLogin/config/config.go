package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Сonfig конфигурация
type Сonfig struct {
	AppPath          string
	HTTPPort         string `json:"httpPort"`
	LogLevel         string `json:"LogLevel"`
	SourceLinesInLog *bool  `json:"SourceLinesInLog"`
}

// NewConfig инициализирует конфиг
func NewConfig() (cfg *Сonfig) {
	l := log.New(os.Stderr, "", 0)
	l.Println("logger started")
	cfg = &Сonfig{}
	cfg.AppPath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	file, err := os.Open(filepath.Join(cfg.AppPath, `config.json`))
	if err != nil {
		l.Fatal("can't open config.json")
	}
	defer file.Close()
	buff, err := ioutil.ReadAll(file)
	if err != nil {
		l.Fatal("can't read config.json")
	}
	err = json.Unmarshal(buff, &cfg)
	if err != nil {
		l.Fatal("can't unmarshal config.json")
	}
	return cfg
}
