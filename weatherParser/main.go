package main

import (
	"net"
	"path/filepath"

	pb "github.com/ffo32167/weather/weatherProto"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Структура для промежуточного хранения данных
type weatherResponse struct {
	City      string
	Month     string
	DayNumber string
	TempDay   string
	TempNight string
	Condition string
}

/* todo:
		улучшить работу с ошибками: в методе cacheLoad при ошибке чтения files, err := ioutil.ReadDir(path)
				создать специальный тип ошибки(путь не найден), и если именно эта ошибка происходит, то
				создавать дерево папок аналогично функции cacheMonthWrite
		реализовать shutdown и интерцепторы для grpc и http
		разбить на пакеты?
-----------------------------------------------------------------------------------------------------------
		добавить модуль сборник типовых действий не вошедших сюда - работа с каналами, контекстом и т.д.
		80% code test coverage
*/

func main() {
	cfg := newConfig()
	log.Info("Reading Cache")
	wmc := newWeatherCache()
	path := filepath.Join(cfg.appPath, `cache`, `yandex`, `russia`)
	wmc.cacheLoad(path)
	log.Info("GRPC service starting up...")
	lis, err := net.Listen("tcp", cfg.GrpcPort)
	if err != nil {
		log.WithFields(logrus.Fields{"err": err}).Fatal("can't listen port")
	}
	serverGRPC := grpc.NewServer()
	pb.RegisterWeatherParserServer(serverGRPC, &grpcServer{config: cfg, wmc: &wmc})
	if err := serverGRPC.Serve(lis); err != nil {
		log.WithFields(logrus.Fields{"err": err}).Fatal("can't start grpc Server")
	}
}
