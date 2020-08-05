package grpcservice

import (
	ch "github.com/ffo32167/weather/weatherParser/cache"
	c "github.com/ffo32167/weather/weatherParser/config"
	w "github.com/ffo32167/weather/weatherParser/weatherresponse"
	pb "github.com/ffo32167/weather/weatherProto"
)

var (
	memoryCache = ch.NewWeatherCache()

	param = &pb.WeatherParams{
		MonthsNumbers: []int32{1, 1},
		Cities:        []string{"Moscow", "Volgodonsk"},
		Site:          "yandex",
		Months:        []string{"january", "january"},
		Year:          "2018",
		ReplyFormat:   "csv",
	}

	weatherDataGrabResult [][]w.WeatherResponse = [][]w.WeatherResponse{
		[]w.WeatherResponse{
			w.WeatherResponse{City: "Moscow", Month: "january", DayNumber: "1", TempDay: "−2°", TempNight: "−6°", Condition: "Облачно и слабый снег"},
			w.WeatherResponse{City: "Moscow", Month: "january", DayNumber: "2", TempDay: "−5°", TempNight: "−6°", Condition: "Облачно"},
			w.WeatherResponse{City: "Moscow", Month: "january", DayNumber: "3", TempDay: "−7°", TempNight: "−9°", Condition: "Облачно и слабый снег"},
			w.WeatherResponse{City: "Moscow", Month: "january", DayNumber: "4", TempDay: "−8°", TempNight: "−9°", Condition: "Облачно и слабый снег"},
			w.WeatherResponse{City: "Moscow", Month: "january", DayNumber: "5", TempDay: "−8°", TempNight: "−11°", Condition: "Облачно"},
			w.WeatherResponse{City: "Moscow", Month: "january", DayNumber: "6", TempDay: "−10°", TempNight: "−11°", Condition: "Облачно"},
			w.WeatherResponse{City: "Moscow", Month: "january", DayNumber: "7", TempDay: "−10°", TempNight: "−11°", Condition: "Облачно"},
			w.WeatherResponse{City: "Moscow", Month: "january", DayNumber: "8", TempDay: "−9°", TempNight: "−9°", Condition: "Облачно"},
			w.WeatherResponse{City: "Moscow", Month: "january", DayNumber: "9", TempDay: "−8°", TempNight: "−9°", Condition: "Облачно и слабый снег"},
			w.WeatherResponse{City: "Moscow", Month: "january", DayNumber: "10", TempDay: "−7°", TempNight: "−8°", Condition: "Облачно и слабый снег"},
			w.WeatherResponse{City: "Moscow", Month: "january", DayNumber: "11", TempDay: "−6°", TempNight: "−7°", Condition: "Облачно и слабый снег"},
			w.WeatherResponse{City: "Moscow", Month: "january", DayNumber: "12", TempDay: "−5°", TempNight: "−5°", Condition: "Облачно и слабый снег"},
			w.WeatherResponse{City: "Moscow", Month: "january", DayNumber: "13", TempDay: "−3°", TempNight: "−4°", Condition: "Облачно и слабый снег"},
			w.WeatherResponse{City: "Moscow", Month: "january", DayNumber: "14", TempDay: "−3°", TempNight: "−5°", Condition: "Облачно"},
			w.WeatherResponse{City: "Moscow", Month: "january", DayNumber: "15", TempDay: "−4°", TempNight: "−6°", Condition: "Облачно и слабый снег"},
			w.WeatherResponse{City: "Moscow", Month: "january", DayNumber: "16", TempDay: "−5°", TempNight: "−7°", Condition: "Облачно и слабый снег"},
			w.WeatherResponse{City: "Moscow", Month: "january", DayNumber: "17", TempDay: "−7°", TempNight: "−9°", Condition: "Облачно"},
			w.WeatherResponse{City: "Moscow", Month: "january", DayNumber: "18", TempDay: "−8°", TempNight: "−10°", Condition: "Облачно и слабый снег"},
			w.WeatherResponse{City: "Moscow", Month: "january", DayNumber: "19", TempDay: "−8°", TempNight: "−10°", Condition: "Облачно"},
			w.WeatherResponse{City: "Moscow", Month: "january", DayNumber: "20", TempDay: "−8°", TempNight: "−10°", Condition: "Ясно"},
			w.WeatherResponse{City: "Moscow", Month: "january", DayNumber: "21", TempDay: "−9°", TempNight: "−11°", Condition: "Облачно"},
			w.WeatherResponse{City: "Moscow", Month: "january", DayNumber: "22", TempDay: "−9°", TempNight: "−9°", Condition: "Облачно и слабый снег"},
			w.WeatherResponse{City: "Moscow", Month: "january", DayNumber: "23", TempDay: "−8°", TempNight: "−9°", Condition: "Облачно"},
			w.WeatherResponse{City: "Moscow", Month: "january", DayNumber: "24", TempDay: "−8°", TempNight: "−10°", Condition: "Облачно"},
			w.WeatherResponse{City: "Moscow", Month: "january", DayNumber: "25", TempDay: "−9°", TempNight: "−12°", Condition: "Облачно"},
			w.WeatherResponse{City: "Moscow", Month: "january", DayNumber: "26", TempDay: "−10°", TempNight: "−11°", Condition: "Облачно и слабый снег"},
			w.WeatherResponse{City: "Moscow", Month: "january", DayNumber: "27", TempDay: "−8°", TempNight: "−9°", Condition: "Облачно"},
			w.WeatherResponse{City: "Moscow", Month: "january", DayNumber: "28", TempDay: "−7°", TempNight: "−9°", Condition: "Ясно"},
			w.WeatherResponse{City: "Moscow", Month: "january", DayNumber: "29", TempDay: "−7°", TempNight: "−9°", Condition: "Облачно и слабый снег"},
			w.WeatherResponse{City: "Moscow", Month: "january", DayNumber: "30", TempDay: "−7°", TempNight: "−8°", Condition: "Облачно"},
			w.WeatherResponse{City: "Moscow", Month: "january", DayNumber: "31", TempDay: "−6°", TempNight: "−8°", Condition: "Облачно и слабый снег"},
		},
		[]w.WeatherResponse{
			w.WeatherResponse{City: "Volgodonsk", Month: "january", DayNumber: "1", TempDay: "−3°", TempNight: "−6°", Condition: "Облачно"},
			w.WeatherResponse{City: "Volgodonsk", Month: "january", DayNumber: "2", TempDay: "−1°", TempNight: "−3°", Condition: "Облачно"},
			w.WeatherResponse{City: "Volgodonsk", Month: "january", DayNumber: "3", TempDay: "−1°", TempNight: "−4°", Condition: "Облачно и слабый снег"},
			w.WeatherResponse{City: "Volgodonsk", Month: "january", DayNumber: "4", TempDay: "−4°", TempNight: "−6°", Condition: "Облачно"},
			w.WeatherResponse{City: "Volgodonsk", Month: "january", DayNumber: "5", TempDay: "−2°", TempNight: "−5°", Condition: "Облачно"},
			w.WeatherResponse{City: "Volgodonsk", Month: "january", DayNumber: "6", TempDay: "−3°", TempNight: "−4°", Condition: "Облачно и слабый снег"},
			w.WeatherResponse{City: "Volgodonsk", Month: "january", DayNumber: "7", TempDay: "−3°", TempNight: "−6°", Condition: "Облачно и слабый снег"},
			w.WeatherResponse{City: "Volgodonsk", Month: "january", DayNumber: "8", TempDay: "−4°", TempNight: "−5°", Condition: "Облачно"},
			w.WeatherResponse{City: "Volgodonsk", Month: "january", DayNumber: "9", TempDay: "−3°", TempNight: "−5°", Condition: "Облачно"},
			w.WeatherResponse{City: "Volgodonsk", Month: "january", DayNumber: "10", TempDay: "−2°", TempNight: "−3°", Condition: "Облачно"},
			w.WeatherResponse{City: "Volgodonsk", Month: "january", DayNumber: "11", TempDay: "−2°", TempNight: "−4°", Condition: "Облачно"},
			w.WeatherResponse{City: "Volgodonsk", Month: "january", DayNumber: "12", TempDay: "−1°", TempNight: "−2°", Condition: "Облачно"},
			w.WeatherResponse{City: "Volgodonsk", Month: "january", DayNumber: "13", TempDay: "+1°", TempNight: "−1°", Condition: "Облачно"},
			w.WeatherResponse{City: "Volgodonsk", Month: "january", DayNumber: "14", TempDay: "0°", TempNight: "−2°", Condition: "Облачно и слабый снег"},
			w.WeatherResponse{City: "Volgodonsk", Month: "january", DayNumber: "15", TempDay: "0°", TempNight: "−1°", Condition: "Облачно"},
			w.WeatherResponse{City: "Volgodonsk", Month: "january", DayNumber: "16", TempDay: "0°", TempNight: "−1°", Condition: "Облачно"},
			w.WeatherResponse{City: "Volgodonsk", Month: "january", DayNumber: "17", TempDay: "−1°", TempNight: "−3°", Condition: "Облачно"},
			w.WeatherResponse{City: "Volgodonsk", Month: "january", DayNumber: "18", TempDay: "−2°", TempNight: "−4°", Condition: "Облачно"},
			w.WeatherResponse{City: "Volgodonsk", Month: "january", DayNumber: "19", TempDay: "−2°", TempNight: "−5°", Condition: "Облачно"},
			w.WeatherResponse{City: "Volgodonsk", Month: "january", DayNumber: "20", TempDay: "−4°", TempNight: "−5°", Condition: "Облачно"},
			w.WeatherResponse{City: "Volgodonsk", Month: "january", DayNumber: "21", TempDay: "−3°", TempNight: "−4°", Condition: "Облачно"},
			w.WeatherResponse{City: "Volgodonsk", Month: "january", DayNumber: "22", TempDay: "−2°", TempNight: "−3°", Condition: "Облачно и слабый снег"},
			w.WeatherResponse{City: "Volgodonsk", Month: "january", DayNumber: "23", TempDay: "−1°", TempNight: "−3°", Condition: "Ясно"},
			w.WeatherResponse{City: "Volgodonsk", Month: "january", DayNumber: "24", TempDay: "−3°", TempNight: "−6°", Condition: "Облачно и слабый снег"},
			w.WeatherResponse{City: "Volgodonsk", Month: "january", DayNumber: "25", TempDay: "−5°", TempNight: "−6°", Condition: "Облачно"},
			w.WeatherResponse{City: "Volgodonsk", Month: "january", DayNumber: "26", TempDay: "−5°", TempNight: "−8°", Condition: "Облачно и слабый снег"},
			w.WeatherResponse{City: "Volgodonsk", Month: "january", DayNumber: "27", TempDay: "−7°", TempNight: "−7°", Condition: "Облачно"},
			w.WeatherResponse{City: "Volgodonsk", Month: "january", DayNumber: "28", TempDay: "−5°", TempNight: "−6°", Condition: "Облачно и слабый снег"},
			w.WeatherResponse{City: "Volgodonsk", Month: "january", DayNumber: "29", TempDay: "−6°", TempNight: "−7°", Condition: "Облачно"},
			w.WeatherResponse{City: "Volgodonsk", Month: "january", DayNumber: "30", TempDay: "−6°", TempNight: "−7°", Condition: "Облачно"},
			w.WeatherResponse{City: "Volgodonsk", Month: "january", DayNumber: "31", TempDay: "−5°", TempNight: "−7°", Condition: "Облачно"},
		},
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
