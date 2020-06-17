package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	pb "github.com/ffo32167/weather/weatherProto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Вызвать удаленную процедуру
func getWeather(cities []string, months []int32, site string) []byte {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
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

func pageWeatherGet(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	untypedUsername := session.Values["user"]
	username, ok := untypedUsername.(string)
	if !ok {
		fmt.Println("cannot assert type for username")
	}
	fmt.Println("pageWeatherGet, values user:", username)
	if err := templates.ExecuteTemplate(w, "index.html", struct{ Login string }{Login: username}); err != nil {
		log.Error("can't parse index.html:", err)
	}
}

// Обработчик
func pageWeatherPost(w http.ResponseWriter, r *http.Request) {
	var filename = "Comparison.csv"
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
	// Вызвать grpc
	// todo сделать обработку ошибок/пустых строк для FormValue
	data := getWeather(
		strings.Split(r.FormValue("cities"), ", "),
		monthsNumbers,
		r.FormValue("Site"),
	)
	// Отдать файл через браузер
	fmt.Println("data size", len(data))
	if len(data) < 100 {
		log.WithFields(logrus.Fields{"data size": len(data), "cities": r.FormValue("cities"), "monthStart": r.FormValue("monthStart"), "monthEnd": r.FormValue("monthEnd")}).Error("failed to get data")
	}
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Write(data)
}
