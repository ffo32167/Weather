package siteparse

import (
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
	c "github.com/ffo32167/weather/weatherParser/config"
	w "github.com/ffo32167/weather/weatherParser/weatherresponse"
	"github.com/sirupsen/logrus"
)

type yandex struct {
	address string
}

//	или https://yandex.ru/pogoda/moscow/month/january
func (y yandex) CreateDataPath(country, city, month, year string) (url string) {
	return (y.address + city + "/month/" + month)
}

//	Распарсить информацию из данных источника
func (yandex) SiteParse(dataSource io.Reader, city string, month string, config c.Config) (data []w.WeatherResponse) {
	var (
		day     w.WeatherResponse
		doWrite bool
	)
	doc, err := goquery.NewDocumentFromReader(dataSource)
	if err != nil {
		logrus.Error("can't parse page as HTML")
	}
	doc.Find(config.YandexWeatherSection).Each(func(i int, s *goquery.Selection) {
		day.City = city
		day.DayNumber = s.Find(config.YandexWeatherDayNumber).Text()
		day.Month = month
		temp := s.Find(config.YandexWeatherTemp).Text()
		tempSplit := strings.SplitAfter(temp, "°")
		if len(temp) > 1 {
			day.TempDay = tempSplit[0]
			day.TempNight = tempSplit[1]
		}
		condition, _ := s.Find(config.YandexWeatherCondition).Attr(config.YandexWeatherConditionAttr)
		//	извлекаем осадки/облачность, т.к в Яндексе нет текстового поля
		if len(condition) > 0 {
			//	берем путь к картинке погоды, извлекаем имя картинки
			picName := condition[strings.LastIndex(condition, "/")+1 : len(condition)-4]
			//	сравниваем имя картинки с мапой соответствия
			day.Condition = config.YandexWeatherMap[picName]
			//	если не находим, то оставляем название картинки
			if day.Condition == "" {
				day.Condition = picName
			}
		}
		// т.к. помимо нужного месяца яндекс дописывает последние/первые числа других месяцев,
		// то первого числа каждого месяца меняем своё желание записывать данные
		if day.DayNumber == "1" {
			doWrite = !doWrite
		}
		if doWrite {
			data = append(data, day)
		}
	})
	return data
}
