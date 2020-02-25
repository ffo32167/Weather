package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

// config это настройки
type config struct {
	appPath  string
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

//	Выбрать источник данных
func (config config) chooseSiteParser(site string) (source siteParser) {
	switch site {
	case "worldweather":
		source = worldWeather{address: config.WorldWeatherAddress}
	default:
		source = yandex{address: config.YandexWeatherAddress}
	}
	return
}

//	Загрузить конфигурацию, настроить логи
func newConfig() (cfg config) {
	cfg = config{}
	cfg.appPath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	if cfg.appPath == "" {
		log.Fatal("can't get application path")
	}
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
	formatLog(*cfg.SourceLinesInLog)
	setLogLevel(cfg.LogLevel)
	return cfg
}
