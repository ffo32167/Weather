package pageparse

import (
	"io"
	"reflect"
	"testing"

	c "github.com/ffo32167/weather/weatherParser/config"
	w "github.com/ffo32167/weather/weatherParser/weatherdata"
)

func Test_worldWeather_CreateDataPath(t *testing.T) {
	type args struct {
		country string
		city    string
		month   string
		year    string
	}
	tests := []struct {
		name    string
		w       worldWeather
		args    args
		wantURL string
	}{
		{"worldWeather CreateDataPath",
			worldWeather{"https://world-weather.ru/pogoda/"},
			args{"russia", "moscow", "january", "2018"},
			"https://world-weather.ru/pogoda/russia/moscow/january-2018",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotURL := tt.w.CreateDataPath(tt.args.country, tt.args.city, tt.args.month, tt.args.year); gotURL != tt.wantURL {
				t.Errorf("worldWeather.CreateDataPath() = %v, want %v", gotURL, tt.wantURL)
			}
		})
	}
}

func Test_worldWeather_SiteParse(t *testing.T) {
	// к сожалению, пока что хардкод :(
	type args struct {
		dataSource io.Reader
		city       string
		month      string
		config     c.Config
	}
	tests := []struct {
		name     string
		w        worldWeather
		args     args
		wantData []w.DayWeather
	}{
		{
			"worldWeather siteParse",
			worldWeather{"https://world-weather.ru/pogoda/"},
			args{worldWeatherMoscowJanuary2018Page, "Moscow", "january", cfg},
			worldWeatherMoscowJanuary2018JSON,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotData := tt.w.SiteParse(tt.args.dataSource, tt.args.city, tt.args.month, tt.args.config); !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("worldWeather.SiteParse() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}
