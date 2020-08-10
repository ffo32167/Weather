package pageparse

import (
	"io"
	"reflect"
	"testing"

	c "github.com/ffo32167/weather/weatherParser/config"
	w "github.com/ffo32167/weather/weatherParser/weatherdata"
)

func Test_yandex_CreateDataPath(t *testing.T) {
	type args struct {
		country string
		city    string
		month   string
		year    string
	}
	tests := []struct {
		name    string
		y       yandex
		args    args
		wantURL string
	}{
		{
			"yandexCreateDataPath",
			yandex{"https://yandex.ru/pogoda/"},
			args{"", "moscow", "january", ""},
			"https://yandex.ru/pogoda/moscow/month/january",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotURL := tt.y.CreateDataPath(tt.args.country, tt.args.city, tt.args.month, tt.args.year); gotURL != tt.wantURL {
				t.Errorf("yandex.CreateDataPath() = %v, want %v", gotURL, tt.wantURL)
			}
		})
	}
}

func Test_yandex_SiteParse(t *testing.T) {
	type args struct {
		dataSource io.Reader
		city       string
		month      string
		config     c.Config
	}
	tests := []struct {
		name     string
		y        yandex
		args     args
		wantData []w.DayWeather
	}{
		{
			"yandex siteParse",
			yandex{"https://yandex.ru/pogoda/"},
			args{yandexMoscowJanuary2018Page, "Moscow", "january", cfg},
			yandexMoscowJanuaryJSON,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotData := tt.y.SiteParse(tt.args.dataSource, tt.args.city, tt.args.month, tt.args.config); !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("yandex.SiteParse() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}
