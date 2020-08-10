package processer

import (
	"bytes"
	"io"
	"net/http"
	"time"

	ch "github.com/ffo32167/weather/weatherParser/cache"
	c "github.com/ffo32167/weather/weatherParser/config"
	e "github.com/ffo32167/weather/weatherParser/encode"
	s "github.com/ffo32167/weather/weatherParser/pageparse"
	w "github.com/ffo32167/weather/weatherParser/weatherdata"

	"github.com/sirupsen/logrus"
)

// ProcessRequest обрабатывает запрос
// из параметров, конфига и кэша делает конечный результат
func ProcessRequest(params w.WeatherParams, config *c.Config, cacher ch.Cacher) (bytes.Buffer, string) {
	siteParser := s.ChooseSiteParser(params.Site, config)
	data := weatherDataGet(siteParser, params, config, cacher)
	encoder := e.ChooseEncoder(params.ReplyFormat)
	return encoder.Encode(data, params.Cities)
}

// weatherDataGet получает данные путём рассматривания кэша или выбранного сайта
func weatherDataGet(sp s.SiteParser, params w.WeatherParams, config *c.Config, cacher ch.Cacher) (wr [][]w.DayWeather) {
	var (
		cityWeather []w.DayWeather
		country     string = "russia"
	)
	//Перевести месяцы в нужный вид
	params.Months = monthsParse(params.MonthsNumbers)
	logrus.WithFields(logrus.Fields{"params": params, "config": config}).Debug("weatherDataGrab parameters")
	for _, city := range params.Cities {
		for _, month := range params.Months {
			// Составить путь к данным в кэше
			path := cacher.Path(config.AppPath, params.Site, country, city, month, params.Year)
			// Проверить есть ли кэш, взять данные из него и перейти к следующему
			monthWeather, _ := cacher.MonthRead(path)
			logrus.WithFields(logrus.Fields{"len(monthWeather)": len(monthWeather)}).Debug("weatherDataGrab parameters")
			if len(monthWeather) > 0 {
				cityWeather = append(cityWeather, monthWeather...)
				continue
			}
			// Если кэша нет, то загружать данные из выбранного сайта
			// Создать из параметров строку запроса
			sitePath := sp.CreateDataPath(country, city, month, params.Year)
			// Получить страницу с сайта
			buf := getDataFromSite(sitePath)
			// Прочитать и распарсить страницу
			monthWeather = sp.SiteParse(buf, city, month, *config)
			// Добавить результаты месяца в кэш
			cacher.MonthWrite(path, monthWeather)
			// Добавить результаты месяца в результаты по городу
			cityWeather = append(cityWeather, monthWeather...)
		}
		// Добавить результаты текущего города в общие результаты городов
		wr = append(wr, cityWeather)
		cityWeather = nil
	}
	logrus.WithFields(logrus.Fields{"wResponse len": len(wr)}).Info("weatherDataGrab work completed")
	return wr
}

// monthsParse разворачивает срез месяцев вида []int{11,2}
// в срез []string{"november", "december", "january", "february"}
func monthsParse(monthsNumbers []int32) (monthsNames []string) {
	calendarMonths := [12]string{"january", "february", "march", "april", "may", "june", "july", "august", "september", "october", "november", "december"}
	//	Проверить параметры
	if len(monthsNumbers) != 2 || monthsNumbers[0] < 1 || monthsNumbers[0] > 12 || monthsNumbers[1] < 1 || monthsNumbers[1] > 12 {
		logrus.WithFields(logrus.Fields{"monthsNumbers": monthsNumbers}).Fatal("incorrect interval of months")
	}
	// Объявить как int32 для protobuf
	var i int32
	//	Если месяца по порядку, то вставить недостающее
	if monthsNumbers[1] > monthsNumbers[0] {
		for i = 0; i < monthsNumbers[1]-monthsNumbers[0]+1; i++ {
			monthsNames = append(monthsNames, calendarMonths[monthsNumbers[0]+i-1])
		}
		//	Если не по порядку, то вставить недостающие до конца года и с начала года
	} else if monthsNumbers[0] > monthsNumbers[1] {
		for i = monthsNumbers[0] - 1; i < 12; i++ {
			monthsNames = append(monthsNames, calendarMonths[i])
		}
		for i = 0; i < monthsNumbers[1]; i++ {
			monthsNames = append(monthsNames, calendarMonths[i])
		}
		//	Если месяц один, то его и вставить
	} else if monthsNumbers[0] == monthsNumbers[1] {
		monthsNames = append(monthsNames, calendarMonths[monthsNumbers[1]-1])
	}
	return monthsNames
}

// getDataFromSite получает страницу для парсинга
func getDataFromSite(sitePath string) io.Reader {
	client := http.Client{
		Timeout: 1 * time.Second,
	}
	resp, err := client.Get(sitePath)
	if err != nil {
		logrus.WithFields(logrus.Fields{"sitePath": sitePath}).Error("can't complete request to site")
	}
	if resp.StatusCode != http.StatusOK {
		logrus.WithFields(logrus.Fields{"sitePath": sitePath, "status code": resp.StatusCode}).Error("weather page not found")
	}
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		logrus.Error("cant copy response:", sitePath)
	}
	return buf
}
