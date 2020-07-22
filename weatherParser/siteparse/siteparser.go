package siteparse

import (
	"io"

	c "github.com/ffo32167/weather/weatherParser/config"
	w "github.com/ffo32167/weather/weatherParser/weatherresponse"
)

// SiteParser Интерфейс для получения данных с сайтов(worldWeather/yandexWeather)
type SiteParser interface {
	CreateDataPath(country, city, month, year string) (address string)
	SiteParse(source io.Reader, city string, month string, config c.Config) []w.WeatherResponse
}

// ChooseSiteParser Выбирает источник данных
func ChooseSiteParser(site string, config *c.Config) (source SiteParser) {
	switch site {
	case "worldweather":
		source = worldWeather{address: config.WorldWeatherAddress}
	default:
		source = yandex{address: config.YandexWeatherAddress}
	}
	return
}
