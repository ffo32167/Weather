package cache

import (
	"path/filepath"
	"reflect"
	"testing"

	w "github.com/ffo32167/weather/weatherParser/weatherdata"
)

var (
	pathMscJan    = `appPath\cache\yandex\russia\moscow_january_2018.json`
	weatherMscJan = []w.DayWeather{
		w.DayWeather{City: "Moscow", Month: "january", DayNumber: "1", TempDay: "−2°", TempNight: "−6°", Condition: "Облачно и слабый снег"},
		w.DayWeather{City: "Moscow", Month: "january", DayNumber: "2", TempDay: "−5°", TempNight: "−6°", Condition: "Облачно"},
		w.DayWeather{City: "Moscow", Month: "january", DayNumber: "3", TempDay: "−7°", TempNight: "−9°", Condition: "Облачно и слабый снег"},
		w.DayWeather{City: "Moscow", Month: "january", DayNumber: "4", TempDay: "−8°", TempNight: "−9°", Condition: "Облачно и слабый снег"},
		w.DayWeather{City: "Moscow", Month: "january", DayNumber: "5", TempDay: "−8°", TempNight: "−11°", Condition: "Облачно"},
		w.DayWeather{City: "Moscow", Month: "january", DayNumber: "6", TempDay: "−10°", TempNight: "−11°", Condition: "Облачно"},
		w.DayWeather{City: "Moscow", Month: "january", DayNumber: "7", TempDay: "−10°", TempNight: "−11°", Condition: "Облачно"},
		w.DayWeather{City: "Moscow", Month: "january", DayNumber: "8", TempDay: "−9°", TempNight: "−9°", Condition: "Облачно"},
		w.DayWeather{City: "Moscow", Month: "january", DayNumber: "9", TempDay: "−8°", TempNight: "−9°", Condition: "Облачно и слабый снег"},
		w.DayWeather{City: "Moscow", Month: "january", DayNumber: "10", TempDay: "−7°", TempNight: "−8°", Condition: "Облачно и слабый снег"},
		w.DayWeather{City: "Moscow", Month: "january", DayNumber: "11", TempDay: "−6°", TempNight: "−7°", Condition: "Облачно и слабый снег"},
		w.DayWeather{City: "Moscow", Month: "january", DayNumber: "12", TempDay: "−5°", TempNight: "−5°", Condition: "Облачно и слабый снег"},
		w.DayWeather{City: "Moscow", Month: "january", DayNumber: "13", TempDay: "−3°", TempNight: "−4°", Condition: "Облачно и слабый снег"},
		w.DayWeather{City: "Moscow", Month: "january", DayNumber: "14", TempDay: "−3°", TempNight: "−5°", Condition: "Облачно"},
		w.DayWeather{City: "Moscow", Month: "january", DayNumber: "15", TempDay: "−4°", TempNight: "−6°", Condition: "Облачно и слабый снег"},
		w.DayWeather{City: "Moscow", Month: "january", DayNumber: "16", TempDay: "−5°", TempNight: "−7°", Condition: "Облачно и слабый снег"},
		w.DayWeather{City: "Moscow", Month: "january", DayNumber: "17", TempDay: "−7°", TempNight: "−9°", Condition: "Облачно"},
		w.DayWeather{City: "Moscow", Month: "january", DayNumber: "18", TempDay: "−8°", TempNight: "−10°", Condition: "Облачно и слабый снег"},
		w.DayWeather{City: "Moscow", Month: "january", DayNumber: "19", TempDay: "−8°", TempNight: "−10°", Condition: "Облачно"},
		w.DayWeather{City: "Moscow", Month: "january", DayNumber: "20", TempDay: "−8°", TempNight: "−10°", Condition: "Ясно"},
		w.DayWeather{City: "Moscow", Month: "january", DayNumber: "21", TempDay: "−9°", TempNight: "−11°", Condition: "Облачно"},
		w.DayWeather{City: "Moscow", Month: "january", DayNumber: "22", TempDay: "−9°", TempNight: "−9°", Condition: "Облачно и слабый снег"},
		w.DayWeather{City: "Moscow", Month: "january", DayNumber: "23", TempDay: "−8°", TempNight: "−9°", Condition: "Облачно"},
		w.DayWeather{City: "Moscow", Month: "january", DayNumber: "24", TempDay: "−8°", TempNight: "−10°", Condition: "Облачно"},
		w.DayWeather{City: "Moscow", Month: "january", DayNumber: "25", TempDay: "−9°", TempNight: "−12°", Condition: "Облачно"},
		w.DayWeather{City: "Moscow", Month: "january", DayNumber: "26", TempDay: "−10°", TempNight: "−11°", Condition: "Облачно и слабый снег"},
		w.DayWeather{City: "Moscow", Month: "january", DayNumber: "27", TempDay: "−8°", TempNight: "−9°", Condition: "Облачно"},
		w.DayWeather{City: "Moscow", Month: "january", DayNumber: "28", TempDay: "−7°", TempNight: "−9°", Condition: "Ясно"},
		w.DayWeather{City: "Moscow", Month: "january", DayNumber: "29", TempDay: "−7°", TempNight: "−9°", Condition: "Облачно и слабый снег"},
		w.DayWeather{City: "Moscow", Month: "january", DayNumber: "30", TempDay: "−7°", TempNight: "−8°", Condition: "Облачно"},
		w.DayWeather{City: "Moscow", Month: "january", DayNumber: "31", TempDay: "−6°", TempNight: "−8°", Condition: "Облачно и слабый снег"},
	}
)

func TestWeatherMemCache_MonthRead(t *testing.T) {
	wmc := NewWeatherMemCache()
	wmc.cache[pathMscJan] = weatherMscJan
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		wmc     *WeatherMemCache
		args    args
		wantWr  []w.DayWeather
		wantErr bool
	}{
		{
			"MonthRead",
			wmc,
			args{pathMscJan},
			weatherMscJan,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWr, err := tt.wmc.MonthRead(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("WeatherMemCache.MonthRead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotWr, tt.wantWr) {
				t.Errorf("WeatherMemCache.MonthRead() = %v, want %v", gotWr, tt.wantWr)
			}
		})
	}
}

func TestWeatherMemCache_Path(t *testing.T) {
	wmc := NewWeatherMemCache()
	type args struct {
		pathParts []string
	}
	tests := []struct {
		name string
		wmc  *WeatherMemCache
		args args
		want string
	}{
		{
			"testPath",
			wmc,
			args{[]string{"appPath", "yandex", "russia", "moscow", "january", "2018"}},
			filepath.Join("appPath", "cache", "yandex", "russia", "moscow_january_2018.json"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.wmc.Path(tt.args.pathParts...); got != tt.want {
				t.Errorf("WeatherMemCache.Path() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeatherMemCache_monthStore(t *testing.T) {
	wmc := NewWeatherMemCache()
	type args struct {
		path string
		wr   []w.DayWeather
	}
	tests := []struct {
		name string
		wmc  *WeatherMemCache
		args args
	}{
		{"monthStore",
			wmc,
			args{
				pathMscJan,
				weatherMscJan,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wmc.monthStore(tt.args.path, tt.args.wr)
		})
	}
}

func Test_path(t *testing.T) {
	type args struct {
		pathParts []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"testPath",
			args{[]string{"appPath", "yandex", "russia", "moscow", "january", "2018"}},
			filepath.Join("appPath", "cache", "yandex", "russia", "moscow_january_2018.json"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := path(tt.args.pathParts...); got != tt.want {
				t.Errorf("path() = %v, want %v", got, tt.want)
			}
		})
	}
}
