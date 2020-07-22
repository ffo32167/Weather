package siteparse

import (
	"io"

	"github.com/PuerkitoBio/goquery"
	c "github.com/ffo32167/weather/weatherParser/config"
	w "github.com/ffo32167/weather/weatherParser/weatherresponse"
	"github.com/sirupsen/logrus"
)

// worldWeather содержит информацию для парсинга сайта worldWeather
type worldWeather struct {
	address string
}

// CreateDataPath cоздаёт путь к нужной странице вида:
// https://world-weather.ru/pogoda/russia/moscow/january-2018/
func (w worldWeather) CreateDataPath(country, city, month, year string) (url string) {
	return (w.address + country + "/" + city + "/" + month + "-" + year)
}

func (worldWeather) SiteParse(dataSource io.Reader, city string, month string, config c.Config) (data []w.WeatherResponse) {
	var day w.WeatherResponse
	doc, err := goquery.NewDocumentFromReader(dataSource)
	if err != nil {
		logrus.Error("can't parse page as HTML")
	}
	doc.Find(config.WorldWeatherSection).Each(func(i int, s *goquery.Selection) {
		day.City = city
		day.DayNumber = s.Find(config.WorldWeatherDayNumber).Text()
		day.Month = month
		day.TempDay = s.Find(config.WorldWeatherTempDay).Text()
		day.TempNight = s.Find(config.WorldWeatherTempNight).Text()
		day.Condition, _ = s.Find(config.WorldWeatherCondition).Attr(config.WorldWeatherConditionAttr)
		data = append(data, day)
	})
	return data
}
