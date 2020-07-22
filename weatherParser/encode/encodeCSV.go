package encode

import (
	"bytes"
	enccsv "encoding/csv"

	w "github.com/ffo32167/weather/weatherParser/weatherresponse"
	"github.com/sirupsen/logrus"
)

// CSV упаковывает файл в csv формат
type csv struct {
	formatName string
}

// NewCSV создаёт CSV
func newCSV() csv {
	return csv{".csv"}
}

// Encode перекодирует данные из среза структур в CSV-файл лежащий в bytes.Buffer
func (c csv) Encode(data [][]w.WeatherResponse, cities []string) (encodedData bytes.Buffer, format string) {
	logrus.WithFields(logrus.Fields{"cities": cities}).Info("start encoding")
	// Создаем заголовки для csv
	cities = append([]string{"days"}, cities...)
	w := enccsv.NewWriter(&encodedData)
	w.Write(cities)
	oneLine := make([]string, 0)
	// Записываем данные построчно
	for i := range data[0] {
		oneLine = append(oneLine, data[0][i].DayNumber+" "+data[0][i].Month)
		for j := range data {
			//	Добавляем только при соответствии дня месяца
			if data[j][i].DayNumber == data[0][i].DayNumber && data[j][i].Month == data[0][i].Month {
				oneLine = append(oneLine, data[j][i].TempDay+" "+data[j][i].TempNight+" "+data[j][i].Condition)
			}
		}
		if err := w.Write(oneLine); err != nil {
			logrus.WithFields(logrus.Fields{"err": err, "with line:": oneLine}).Error("error writing record to csv")
		}
		oneLine = nil
	}
	w.Flush()
	if w.Error() != nil {
		logrus.Error("error flushing last record to csv")
	}
	return encodedData, c.formatName
}
