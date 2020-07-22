package grpcservice

import (
	"context"
	"net"
	"net/http"
	"time"

	ch "github.com/ffo32167/weather/weatherParser/cache"
	c "github.com/ffo32167/weather/weatherParser/config"
	e "github.com/ffo32167/weather/weatherParser/encode"
	s "github.com/ffo32167/weather/weatherParser/siteparse"
	w "github.com/ffo32167/weather/weatherParser/weatherresponse"
	pb "github.com/ffo32167/weather/weatherProto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// GrpcServer реализует сервер grpc
type GrpcServer struct {
	pb.UnimplementedWeatherParserServer
	config *c.Config
	wmc    *ch.WeatherMemCache
}

// ServerStart запускает сервер GRPC
func ServerStart(cfg *c.Config, wmc *ch.WeatherMemCache) {
	lis, err := net.Listen("tcp", cfg.GrpcPort)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Fatal("can't listen port")
	}
	serverGRPC := grpc.NewServer()
	serv := NewGrpcServer(cfg, wmc)
	pb.RegisterWeatherParserServer(serverGRPC, serv)
	if err := serverGRPC.Serve(lis); err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Fatal("can't start grpc Server")
	}
	logrus.Info("server started")
}

// NewGrpcServer создаёт сервер из конфига и кэша
func NewGrpcServer(cfg *c.Config, wmc *ch.WeatherMemCache) (s *GrpcServer) {
	return &GrpcServer{config: cfg, wmc: wmc}
}

// GetWeather обрабатывает grpc запрос
func (server *GrpcServer) GetWeather(ctx context.Context, params *pb.WeatherParams) (*pb.WeatherResponse, error) {
	logrus.WithFields(logrus.Fields{"params": params, "config": server.config}).Debug("grpc params")
	siteParser := s.ChooseSiteParser(params.Site, server.config)
	data := WeatherDataGrab(siteParser, params, server.config, server.wmc)
	encoder := e.ChooseEncoder(params.ReplyFormat)
	buffer, format := encoder.Encode(data, params.Cities)
	logrus.WithFields(logrus.Fields{"params": params, "config": server.config, "result": len(buffer.Bytes())}).Debug("grpc results")
	return &pb.WeatherResponse{ComparisonCSV: buffer.Bytes(), Format: format}, nil
}

// WeatherDataGrab получает данные путём рассматривания кэша или выбранного сайта
func WeatherDataGrab(sp s.SiteParser, params *pb.WeatherParams, config *c.Config, cr ch.CacheReader) (wr [][]w.WeatherResponse) {
	var (
		cityWeather []w.WeatherResponse
		country     string = "russia"
	)
	// перевести месяцы в нужный вид
	params.Months = monthsParse(params.MonthsNumbers)
	logrus.WithFields(logrus.Fields{"params": params, "config": config}).Debug("weatherDataGrab parameters")
	for _, city := range params.Cities {
		for _, month := range params.Months {
			path := ch.Path(config.AppPath, params.Site, country, city, month, params.Year)
			// Проверить есть ли кэш, взять данные из него и перейти к следующему
			monthWeather, _ := cr.CacheRead(path)
			logrus.WithFields(logrus.Fields{"len(monthWeather)": len(monthWeather)}).Debug("weatherDataGrab parameters")
			if len(monthWeather) > 0 {
				cityWeather = append(cityWeather, monthWeather...)
				continue
			}
			// Если кэша нет, то загружать данные из выбранного сайта
			// Создать из параметров строку запроса
			sitePath := sp.CreateDataPath(country, city, month, params.Year)
			client := http.Client{
				Timeout: 1 * time.Second,
			}
			resp, err := client.Get(sitePath)
			if err != nil {
				logrus.WithFields(logrus.Fields{"sitePath": sitePath, "params": params, "config": config, "error": err}).Fatal("can't complete request to site")
			}
			defer resp.Body.Close()
			// Распарсить страницу
			monthWeather = sp.SiteParse(resp.Body, city, month, *config)
			// Добавить результаты месяца в кэш
			cr.CacheMonthWrite(path, monthWeather)
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
	//	если не сделать, то будет mismatched type int(Go) and int32(proto)  :(
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
