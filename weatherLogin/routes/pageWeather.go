package routes

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	s "github.com/ffo32167/weather/weatherLogin/session"
	t "github.com/ffo32167/weather/weatherLogin/templates"
	pb "github.com/ffo32167/weather/weatherProto"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const year = "2018"

// обрабатывает запросы к grpc серверу
type pageWeatherHandler struct {
	grpcConn *grpc.ClientConn
}

func (pwh pageWeatherHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var filename = "Comparison"
	monthsNumbers := make([]int32, 0)
	monthStart, err := strconv.Atoi(r.FormValue("monthStart"))
	if err != nil {
		logrus.WithFields(logrus.Fields{"monthStart": r.FormValue("monthStart")}).Error("can't parse value of monthStart on innerPage")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	monthsNumbers = append(monthsNumbers, int32(monthStart))
	monthEnd, err := strconv.Atoi(r.FormValue("monthEnd"))
	if err != nil {
		logrus.WithFields(logrus.Fields{"monthEnd": r.FormValue("monthEnd")}).Error("can't parse value of monthEnd on innerPage")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	monthsNumbers = append(monthsNumbers, int32(monthEnd))
	// Вызвать grpc
	// todo сделать обработку ошибок/пустых строк для FormValue
	data, format := getWeather(
		strings.Split(r.FormValue("cities"), ", "),
		monthsNumbers,
		r.FormValue("Site"),
		r.FormValue("Format"),
		pwh.grpcConn,
	)
	fullFileName := filename + format
	// Отдать файл через браузер
	logrus.Info("data size ", len(data))
	if len(data) < 100 {
		logrus.WithFields(logrus.Fields{"data size": len(data), "cities": r.FormValue("cities"), "monthStart": r.FormValue("monthStart"), "monthEnd": r.FormValue("monthEnd")}).Error("failed to get data")
	}
	w.Header().Set("Content-Disposition", "attachment; filename="+fullFileName)
	w.Write(data)
}

// Вызвать удаленную процедуру
func getWeather(cities []string, months []int32, site string, replyFormat string, grpcConn *grpc.ClientConn) ([]byte, string) {
	grpcClient := pb.NewWeatherParserClient(grpcConn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := grpcClient.GetWeather(ctx, &pb.WeatherParams{Cities: cities, MonthsNumbers: months, Site: site, Year: year, ReplyFormat: replyFormat})
	if err != nil {
		logrus.Error("grpc error:", err)
	}
	return r.GetComparisonCSV(), r.GetFormat()
}

// Обработчик Get страницы получения погоды
func pageWeatherGet(w http.ResponseWriter, r *http.Request) {
	session, err := s.Store.Get(r, "session")
	if err != nil {
		logrus.Error("can't decode session")
	}
	untypedUsername := session.Values["user"]
	username, ok := untypedUsername.(string)
	if !ok {
		logrus.Error("cannot assert type for username")
	}
	logrus.Info("pageWeatherGet, values user:", username)
	if err := t.Templates.ExecuteTemplate(w, "index.html", struct{ Login string }{Login: username}); err != nil {
		logrus.Error("can't parse index.html:", err)
	}
}

/*
// PageWeatherPost Обработчик Post страницы
func PageWeatherPost(w http.ResponseWriter, r *http.Request) {
	var filename = "Comparison"
	monthsNumbers := make([]int32, 0)
	monthStart, err := strconv.Atoi(r.FormValue("monthStart"))
	if err != nil {
		logrus.WithFields(logrus.Fields{"monthStart": r.FormValue("monthStart")}).Error("can't parse value of monthStart on innerPage")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	monthsNumbers = append(monthsNumbers, int32(monthStart))
	monthEnd, err := strconv.Atoi(r.FormValue("monthEnd"))
	if err != nil {
		logrus.WithFields(logrus.Fields{"monthEnd": r.FormValue("monthEnd")}).Error("can't parse value of monthEnd on innerPage")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	monthsNumbers = append(monthsNumbers, int32(monthEnd))
	// Вызвать grpc
	// todo сделать обработку ошибок/пустых строк для FormValue
	data, format := getWeather(
		strings.Split(r.FormValue("cities"), ", "),
		monthsNumbers,
		r.FormValue("Site"),
		r.FormValue("Format"),
	)
	fullFileName := filename + format
	// Отдать файл через браузер
	logrus.Info("data size", len(data))
	if len(data) < 100 {
		logrus.WithFields(logrus.Fields{"data size": len(data), "cities": r.FormValue("cities"), "monthStart": r.FormValue("monthStart"), "monthEnd": r.FormValue("monthEnd")}).Error("failed to get data")
	}
	w.Header().Set("Content-Disposition", "attachment; filename="+fullFileName)
	w.Write(data)
}
*/
