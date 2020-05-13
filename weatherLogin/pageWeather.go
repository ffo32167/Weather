package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	pb "github.com/ffo32167/weather/weatherProto"
	"google.golang.org/grpc"
)

// пишем только через шаблоны, потому что у шаблонов есть защита от межсайтового скриптинга XSS

// Вызвать удаленную процедуру
func getWeather(cities []string, months []int32, site string) []byte {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Error("can't connect to grpc server:", err)
	}
	defer conn.Close()
	grpcClient := pb.NewWeatherParserClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := grpcClient.GetWeather(ctx, &pb.WeatherParams{Cities: cities, MonthsNumbers: months, Site: site, Year: "2018"})
	if err != nil {
		log.Error("grpc error:", err)
	}
	return r.GetComparisonCSV()
}

func (config *config) pageWeatherGet(w http.ResponseWriter, r *http.Request) {
	// Проверить наличие сессии
	sess, err := checkSession(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Если нет сессии, то отправить за новой
	if sess == nil {
		loginTempl, err := template.New("login").Parse(loginPageTmpl)
		if err != nil {
			log.WithFields(logrus.Fields{"config": config}).Error("can't parse loginPageTmpl template")
		}
		loginTempl.Execute(w, sess)
		// и начинаем всё заново
		return
	}
	// Если только получили страницу, то и отправить страницу
	t, err := template.New("inner").Parse(innerPageTmpl)
	if err != nil {
		log.WithFields(logrus.Fields{"config": config}).Error("can't parse innerPageTmpl template")
	}
	t.Execute(w, sess)
}

// Обработчик
func (config *config) pageWeatherPost(w http.ResponseWriter, r *http.Request) {
	var filename = "Comparison.csv"
	// Если нажали кнопку, то разобрать параметры, выполнить запрос и отправить данные
	monthsNumbers := make([]int32, 0)
	monthStart, err := strconv.Atoi(r.FormValue("monthStart"))
	if err != nil {
		log.WithFields(logrus.Fields{"monthStart": r.FormValue("monthStart")}).Error("can't parse value of monthStart on innerPage")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	monthsNumbers = append(monthsNumbers, int32(monthStart))
	monthEnd, err := strconv.Atoi(r.FormValue("monthEnd"))
	if err != nil {
		log.WithFields(logrus.Fields{"monthEnd": r.FormValue("monthEnd")}).Error("can't parse value of monthEnd on innerPage")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	monthsNumbers = append(monthsNumbers, int32(monthEnd))
	// Получить данные через grpc
	data := getWeather(strings.Split(r.FormValue("cities"), ", "), monthsNumbers, "yandex")
	// Отдать файл через браузер
	// Попробовать mime.TypeByExtension()
	// dataHeader := make([]byte, 512)
	// if len(data) <= 512 {
	// 	copy(dataHeader, data)
	// } else {
	// 	copy(dataHeader, data[:512])
	// }
	fmt.Println("data size", len(data))
	if len(data) < 100 {
		log.WithFields(logrus.Fields{"data size": len(data), "monthStart": r.FormValue("monthStart"), "monthEnd": r.FormValue("monthEnd")}).Error("failed to get data")
	}
	// dataContentType := http.DetectContentType(dataHeader)
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	// w.Header().Set("Content-Type", dataContentType)
	w.Write(data)
}
