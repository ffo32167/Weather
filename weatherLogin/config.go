package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type config struct {
	appPath          string
	HTTPPort         string `json:"httpPort"`
	LogLevel         string `json:"LogLevel"`
	SourceLinesInLog *bool  `json:"SourceLinesInLog"`
}

func newConfig() (cfg *config) {
	cfg = &config{}
	cfg.appPath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	newLog(cfg.appPath)
	file, err := os.Open(filepath.Join(cfg.appPath, `config.json`))
	if err != nil {
		log.Fatal("can't open config.json")
	}
	defer file.Close()
	buff, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("can't read config.json")
	}
	err = json.Unmarshal(buff, &cfg)
	if err != nil {
		log.Fatal("can't unmarshal config.json")
	}
	setLogFormat(*cfg.SourceLinesInLog)
	setLogLevel(cfg.LogLevel)
	return cfg
}
