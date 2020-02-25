package main

import (
	"bytes"
	"encoding/csv"

	"github.com/sirupsen/logrus"
)

//	Перекодировать данные из среза структур в CSV-файл лежащий в bytes.Buffer
func encode(data [][]weatherResponse, cities []string) (encodedData bytes.Buffer) {
	log.WithFields(logrus.Fields{"cities": cities}).Info("start encoding")
	// Создаем заголовки для csv
	cities = append([]string{"days"}, cities...)
	w := csv.NewWriter(&encodedData)
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
			log.WithFields(logrus.Fields{"err": err, "with line:": oneLine}).Error("error writing record to csv")
		}
		oneLine = nil
	}
	w.Flush()
	if w.Error() != nil {
		log.Error("error flushing last record to csv")
	}
	return encodedData
}
