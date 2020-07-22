package cache

import w "github.com/ffo32167/weather/weatherParser/weatherresponse"

// CacheReader создает путь к кешу и сохраняет данные месяца в кеш
type CacheReader interface {
	CacheRead(path string) ([]w.WeatherResponse, error)
	CacheMonthWrite(string, []w.WeatherResponse)
}
