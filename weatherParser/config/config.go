package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Config содержит пути к сайтам, алиасы полей для парсинга и настройки
type Config struct {
	AppPath  string
	LogLevel string `json:"LogLevel"`
	GrpcPort string `json:"GrpcPort"`
	// SourceLinesInLog определяет нужно ли указывать номера строк там где было вызвано логирование
	SourceLinesInLog *bool `json:"SourceLinesInLog"`
	// поля для парсинга данных из Yandex
	YandexWeatherAddress       string            `json:"YandexWeatherAddress"`
	YandexWeatherSection       string            `json:"YandexWeatherSection"`
	YandexWeatherDayNumber     string            `json:"YandexWeatherDayNumber"`
	YandexWeatherTemp          string            `json:"YandexWeatherTemp"`
	YandexWeatherCondition     string            `json:"YandexWeatherCondition"`
	YandexWeatherConditionAttr string            `json:"YandexWeatherConditionAttr"`
	YandexWeatherMap           map[string]string `json:"YandexWeatherMap"`
	// поля для парсинга данных из WorldWeather
	WorldWeatherAddress       string `json:"WorldWeatherAddress"`
	WorldWeatherSection       string `json:"WorldWeatherSection"`
	WorldWeatherDayNumber     string `json:"WorldWeatherDayNumber"`
	WorldWeatherTempDay       string `json:"WorldWeatherTempDay"`
	WorldWeatherTempNight     string `json:"WorldWeatherTempNight"`
	WorldWeatherCondition     string `json:"WorldWeatherCondition"`
	WorldWeatherConditionAttr string `json:"WorldWeatherConditionAttr"`
}

// NewConfig загрузить конфигурацию, настроить логи
func NewConfig() (cfg *Config, err error) {
	cfg = &Config{}
	cfg.AppPath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	if cfg.AppPath == "" {
		return nil, errors.New("can't get application path")
	}
	file, err := os.Open(filepath.Join(cfg.AppPath, `config.json`))
	if err != nil {
		return nil, errors.New("can't open config.json")
	}
	defer file.Close()
	buff, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.New("can't read config.json")
	}
	err = json.Unmarshal(buff, &cfg)
	if err != nil {
		return nil, errors.New("can't unmarshal config.json")
	}
	return cfg, nil
}
