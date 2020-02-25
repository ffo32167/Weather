package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"path/filepath"

	pb "github.com/ffo32167/weather/weatherProto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type cacheReader interface {
	cacheRead(path string) []weatherResponse
	cacheMonthWrite(string, []weatherResponse)
}

// Читает данные из файлов кэша
type app struct {
}

// Запустить как приложение
func runApp(config config) {
	params := loadParams(config.appPath)
	siteParserInterface := config.chooseSiteParser(params.Site)
	log.WithFields(logrus.Fields{"params": params, "config": config}).Debug("runApp parameters")
	app := app{}
	data := weatherDataGrab(siteParserInterface, params, config, app)
	buff := encode(data, params.Cities)
	file, err := os.Create(filepath.Join(config.appPath, `Comparison.csv`))
	defer file.Close()
	if err != nil {
		log.WithFields(logrus.Fields{"params": params, "config": config}).Error("can't create Comparison.csv")
	}
	buff.WriteTo(file)
}

// Реализация сервера grpc
type grpcServer struct {
	pb.UnimplementedWeatherParserServer
	config config
	wmc    *weatherMemCache
}

// Запустить как сервис
func runService(config config) {
	log.Info("Reading Cache")
	wmc := newWeatherCache()
	path := filepath.Join(config.appPath, `cache`, `yandex`, `russia`)
	wmc.cacheLoad(path)
	log.Info("GRPC service starting up...")
	lis, err := net.Listen("tcp", config.GrpcPort)
	if err != nil {
		log.WithFields(logrus.Fields{"err": err}).Fatal("can't listen port")
	}
	serverGRPC := grpc.NewServer()
	pb.RegisterWeatherParserServer(serverGRPC, &grpcServer{config: config, wmc: &wmc})
	if err := serverGRPC.Serve(lis); err != nil {
		log.WithFields(logrus.Fields{"err": err}).Fatal("can't start grpc Server")
	}
}

// Обработка grpc запроса
func (server *grpcServer) GetWeather(ctx context.Context, params *pb.WeatherParams) (*pb.WeatherResponse, error) {
	log.WithFields(logrus.Fields{"params": params, "config": server.config}).Debug("grpc params")
	siteParser := server.config.chooseSiteParser(params.Site)
	data := weatherDataGrab(siteParser, params, server.config, server.wmc)
	buff := encode(data, params.Cities)
	log.WithFields(logrus.Fields{"params": params, "config": server.config, "result": len(buff.Bytes())}).Debug("grpc results")
	return &pb.WeatherResponse{ComparisonCSV: buff.Bytes()}, nil
}

// Получить данные путём рассматривания кэша или выбранного сайта
func weatherDataGrab(sp siteParser, params *pb.WeatherParams, config config, cr cacheReader) (wResponse [][]weatherResponse) {
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
			monthWeather := cr.cacheRead(path)
			if len(monthWeather) > 0 {
				cityWeather = append(cityWeather, monthWeather...)
				continue
			}
			// Если кэша нет, то загружать данные из выбранного сайта
			// Создать из параметров строку запроса
			sitePath := sp.createDataPath(country, city, month, params.Year)
			resp, err := http.Get(sitePath)
			if err != nil {
				log.WithFields(logrus.Fields{"sitePath": sitePath, "params": params, "config": config, "error": err}).Fatal("can't complete request to site")
			}
			// Распарсить страницу
			monthWeather = sp.siteParse(resp.Body, city, month, config)
			// Добавить результаты месяца в кэш
			cr.cacheMonthWrite(path, monthWeather)
			// Добавить результаты месяца в результаты по городу
			cityWeather = append(cityWeather, monthWeather...)
		}
		// Добавить результаты текущего города в общие результаты городов
		wResponse = append(wResponse, cityWeather)
		cityWeather = nil
	}
	log.WithFields(logrus.Fields{"wResponse len": len(wResponse)}).Info("weatherDataGrab work completed")
	return wResponse
}
