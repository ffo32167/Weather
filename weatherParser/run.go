package main

import (
	"context"
	"net/http"
	"time"

	pb "github.com/ffo32167/weather/weatherProto"
	"github.com/sirupsen/logrus"
)

type cacheReader interface {
	cacheRead(path string) ([]weatherResponse, error)
	cacheMonthWrite(string, []weatherResponse)
}

// Реализация сервера grpc
type grpcServer struct {
	pb.UnimplementedWeatherParserServer
	config *config
	wmc    *weatherMemCache
}

// Обработка grpc запроса
func (server *grpcServer) GetWeather(ctx context.Context, params *pb.WeatherParams) (*pb.WeatherResponse, error) {
	log.WithFields(logrus.Fields{"params": params, "config": server.config}).Debug("grpc params")
	siteParser := server.config.chooseSiteParser(params.Site)
	data := weatherDataGrab(siteParser, params, server.config, server.wmc)
	buffer := encode(data, params.Cities)
	log.WithFields(logrus.Fields{"params": params, "config": server.config, "result": len(buffer.Bytes())}).Debug("grpc results")
	return &pb.WeatherResponse{ComparisonCSV: buffer.Bytes()}, nil
}

// Получить данные путём рассматривания кэша или выбранного сайта
func weatherDataGrab(sp siteParser, params *pb.WeatherParams, config *config, cr cacheReader) (wr [][]weatherResponse) {
	var (
		cityWeather []weatherResponse
		country     string = "russia"
	)
	// перевести месяцы в нужный вид
	params.Months = monthsParse(params.MonthsNumbers)
	log.WithFields(logrus.Fields{"params": params, "config": config}).Debug("weatherDataGrab parameters")
	for _, city := range params.Cities {
		for _, month := range params.Months {
			path := cachePath(config.appPath, params.Site, country, city, month, params.Year)
			// Проверить есть ли кэш, взять данные из него и перейти к следующему
			monthWeather, _ := cr.cacheRead(path)
			log.WithFields(logrus.Fields{"len(monthWeather)": len(monthWeather)}).Debug("weatherDataGrab parameters")
			if len(monthWeather) > 0 {
				cityWeather = append(cityWeather, monthWeather...)
				continue
			}
			// Если кэша нет, то загружать данные из выбранного сайта
			// Создать из параметров строку запроса
			sitePath := sp.createDataPath(country, city, month, params.Year)
			client := http.Client{
				Timeout: 1 * time.Second,
			}
			resp, err := client.Get(sitePath)
			if err != nil {
				log.WithFields(logrus.Fields{"sitePath": sitePath, "params": params, "config": config, "error": err}).Fatal("can't complete request to site")
			}
			defer resp.Body.Close()
			// Распарсить страницу
			monthWeather = sp.siteParse(resp.Body, city, month, *config)
			// Добавить результаты месяца в кэш
			cr.cacheMonthWrite(path, monthWeather)
			// Добавить результаты месяца в результаты по городу
			cityWeather = append(cityWeather, monthWeather...)
		}
		// Добавить результаты текущего города в общие результаты городов
		wr = append(wr, cityWeather)
		cityWeather = nil
	}
	log.WithFields(logrus.Fields{"wResponse len": len(wr)}).Info("weatherDataGrab work completed")
	return wr
}
