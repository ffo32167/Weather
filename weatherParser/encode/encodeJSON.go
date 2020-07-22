package encode

import (
	"bytes"
	encjson "encoding/json"

	w "github.com/ffo32167/weather/weatherParser/weatherresponse"
	"github.com/sirupsen/logrus"
)

// JSON упаковывает файл в JSON формат
type json struct {
	formatName string
}

// newJSON создаёт JSON
func newJSON() json {
	return json{".json"}
}

// Encode упаковывает данные в json формат для ответа
func (j json) Encode(dataStruct [][]w.WeatherResponse, cities []string) (encodedData bytes.Buffer, format string) {
	data, err := encjson.Marshal(dataStruct)
	if err != nil {
		logrus.Error("can't encode", dataStruct)
	}
	return *bytes.NewBuffer(data), j.formatName
}
