package encode

import (
	"bytes"

	w "github.com/ffo32167/weather/weatherParser/weatherdata"
)

// Encoder кодирует ответ в различные форматы
type Encoder interface {
	Encode([][]w.DayWeather, []string) (bytes.Buffer, string)
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
