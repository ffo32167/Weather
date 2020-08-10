package processer

import (
	"reflect"
	"testing"

	ch "github.com/ffo32167/weather/weatherParser/cache"
	c "github.com/ffo32167/weather/weatherParser/config"
	s "github.com/ffo32167/weather/weatherParser/pageparse"
	w "github.com/ffo32167/weather/weatherParser/weatherdata"
)

func TestWeatherDataGet(t *testing.T) {
	wmc := ch.NewWeatherMemCache()
	wmc.MonthWrite(pathMscJan, weatherMscJan)
	type args struct {
		sp     s.SiteParser
		params w.WeatherParams
		config *c.Config
		cr     ch.Cacher
	}
	tests := []struct {
		name   string
		args   args
		wantWr [][]w.DayWeather
	}{
		{
			"WeatherDataGrab",
			args{
				s.ChooseSiteParser("yandex", &cfg),
				param,
				&cfg,
				wmc,
			},
			weatherDataGetResult,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotWr := weatherDataGet(tt.args.sp, tt.args.params, tt.args.config, tt.args.cr); !reflect.DeepEqual(gotWr, tt.wantWr) {
				t.Errorf("weatherDataGet() = %v, want %v", gotWr, tt.wantWr)
			}
		})
	}
}

func Test_monthsParse(t *testing.T) {
	type args struct {
		monthsNumbers []int32
	}
	tests := []struct {
		name            string
		args            args
		wantMonthsNames []string
	}{
		{"monthsParse",
			args{[]int32{12, 1}},
			[]string{"december", "january"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMonthsNames := monthsParse(tt.args.monthsNumbers); !reflect.DeepEqual(gotMonthsNames, tt.wantMonthsNames) {
				t.Errorf("monthsParse() = %v, want %v", gotMonthsNames, tt.wantMonthsNames)
			}
		})
	}
}

var (
	memoryCache = ch.NewWeatherMemCache()
	param       = w.WeatherParams{
		MonthsNumbers: []int32{1, 1},
		Cities:        []string{"Moscow"},
		Site:          "yandex",
		Months:        []string{"january", "january"},
		Year:          "2018",
		ReplyFormat:   "csv",
	}

	weatherDataGetResult [][]w.DayWeather = [][]w.DayWeather{
		weatherMscJan,
	}
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

	cfg = c.Config{
		LogLevel:             "debug",
		GrpcPort:             ":50051",
		YandexWeatherAddress: "https://yandex.ru/pogoda/",
		YandexWeatherMap: map[string]string{
			"ovc_sn":   "Облачно и слабый снег",
			"ovc":      "Облачно",
			"skc_d":    "Ясно",
			"ovc_ra":   "Дождь",
			"bkn_ra_d": "Преимущественно ясно и слабый дождь",
		},
		YandexWeatherSection:       ".climate-calendar__cell",
		YandexWeatherDayNumber:     ".climate-calendar-day__day",
		YandexWeatherTemp:          ".climate-calendar-day__temp",
		YandexWeatherCondition:     "img",
		YandexWeatherConditionAttr: "src",
		WorldWeatherAddress:        "https://world-weather.ru/pogoda/",
		WorldWeatherSection:        ".ww-month a",
		WorldWeatherDayNumber:      "div",
		WorldWeatherTempDay:        "span",
		WorldWeatherTempNight:      "p",
		WorldWeatherCondition:      "i",
		WorldWeatherConditionAttr:  "title"}
)
