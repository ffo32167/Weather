package encode

import (
	"bytes"

	w "github.com/ffo32167/weather/weatherParser/weatherresponse"
)

// Encoder кодирует ответ в различные форматы
type Encoder interface {
	Encode([][]w.WeatherResponse, []string) (bytes.Buffer, string)
}

// ChooseEncoder выбирает формат ответа
func ChooseEncoder(format string) (e Encoder) {
	switch format {
	case "JSON":
		e = newJSON()
	default:
		e = newCSV()
	}
	return
}
