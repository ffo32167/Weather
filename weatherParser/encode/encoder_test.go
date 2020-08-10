package encode

import (
	"reflect"
	"testing"

	w "github.com/ffo32167/weather/weatherParser/weatherdata"
)

func TestChooseEncoder(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name  string
		args  args
		wantE Encoder
	}{
		{
			"chooseEncoder",
			args{".sfaasdf"},
			newCSV(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotE := ChooseEncoder(tt.args.format); !reflect.DeepEqual(gotE, tt.wantE) {
				t.Errorf("ChooseEncoder() = %v, want %v", gotE, tt.wantE)
			}
		})
	}
}

var (
	encodeData [][]w.DayWeather = [][]w.DayWeather{
		[]w.DayWeather{
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
		},
		[]w.DayWeather{
			w.DayWeather{City: "Volgodonsk", Month: "january", DayNumber: "1", TempDay: "−3°", TempNight: "−6°", Condition: "Облачно"},
			w.DayWeather{City: "Volgodonsk", Month: "january", DayNumber: "2", TempDay: "−1°", TempNight: "−3°", Condition: "Облачно"},
			w.DayWeather{City: "Volgodonsk", Month: "january", DayNumber: "3", TempDay: "−1°", TempNight: "−4°", Condition: "Облачно и слабый снег"},
			w.DayWeather{City: "Volgodonsk", Month: "january", DayNumber: "4", TempDay: "−4°", TempNight: "−6°", Condition: "Облачно"},
			w.DayWeather{City: "Volgodonsk", Month: "january", DayNumber: "5", TempDay: "−2°", TempNight: "−5°", Condition: "Облачно"},
			w.DayWeather{City: "Volgodonsk", Month: "january", DayNumber: "6", TempDay: "−3°", TempNight: "−4°", Condition: "Облачно и слабый снег"},
			w.DayWeather{City: "Volgodonsk", Month: "january", DayNumber: "7", TempDay: "−3°", TempNight: "−6°", Condition: "Облачно и слабый снег"},
			w.DayWeather{City: "Volgodonsk", Month: "january", DayNumber: "8", TempDay: "−4°", TempNight: "−5°", Condition: "Облачно"},
			w.DayWeather{City: "Volgodonsk", Month: "january", DayNumber: "9", TempDay: "−3°", TempNight: "−5°", Condition: "Облачно"},
			w.DayWeather{City: "Volgodonsk", Month: "january", DayNumber: "10", TempDay: "−2°", TempNight: "−3°", Condition: "Облачно"},
			w.DayWeather{City: "Volgodonsk", Month: "january", DayNumber: "11", TempDay: "−2°", TempNight: "−4°", Condition: "Облачно"},
			w.DayWeather{City: "Volgodonsk", Month: "january", DayNumber: "12", TempDay: "−1°", TempNight: "−2°", Condition: "Облачно"},
			w.DayWeather{City: "Volgodonsk", Month: "january", DayNumber: "13", TempDay: "+1°", TempNight: "−1°", Condition: "Облачно"},
			w.DayWeather{City: "Volgodonsk", Month: "january", DayNumber: "14", TempDay: "0°", TempNight: "−2°", Condition: "Облачно и слабый снег"},
			w.DayWeather{City: "Volgodonsk", Month: "january", DayNumber: "15", TempDay: "0°", TempNight: "−1°", Condition: "Облачно"},
			w.DayWeather{City: "Volgodonsk", Month: "january", DayNumber: "16", TempDay: "0°", TempNight: "−1°", Condition: "Облачно"},
			w.DayWeather{City: "Volgodonsk", Month: "january", DayNumber: "17", TempDay: "−1°", TempNight: "−3°", Condition: "Облачно"},
			w.DayWeather{City: "Volgodonsk", Month: "january", DayNumber: "18", TempDay: "−2°", TempNight: "−4°", Condition: "Облачно"},
			w.DayWeather{City: "Volgodonsk", Month: "january", DayNumber: "19", TempDay: "−2°", TempNight: "−5°", Condition: "Облачно"},
			w.DayWeather{City: "Volgodonsk", Month: "january", DayNumber: "20", TempDay: "−4°", TempNight: "−5°", Condition: "Облачно"},
			w.DayWeather{City: "Volgodonsk", Month: "january", DayNumber: "21", TempDay: "−3°", TempNight: "−4°", Condition: "Облачно"},
			w.DayWeather{City: "Volgodonsk", Month: "january", DayNumber: "22", TempDay: "−2°", TempNight: "−3°", Condition: "Облачно и слабый снег"},
			w.DayWeather{City: "Volgodonsk", Month: "january", DayNumber: "23", TempDay: "−1°", TempNight: "−3°", Condition: "Ясно"},
			w.DayWeather{City: "Volgodonsk", Month: "january", DayNumber: "24", TempDay: "−3°", TempNight: "−6°", Condition: "Облачно и слабый снег"},
			w.DayWeather{City: "Volgodonsk", Month: "january", DayNumber: "25", TempDay: "−5°", TempNight: "−6°", Condition: "Облачно"},
			w.DayWeather{City: "Volgodonsk", Month: "january", DayNumber: "26", TempDay: "−5°", TempNight: "−8°", Condition: "Облачно и слабый снег"},
			w.DayWeather{City: "Volgodonsk", Month: "january", DayNumber: "27", TempDay: "−7°", TempNight: "−7°", Condition: "Облачно"},
			w.DayWeather{City: "Volgodonsk", Month: "january", DayNumber: "28", TempDay: "−5°", TempNight: "−6°", Condition: "Облачно и слабый снег"},
			w.DayWeather{City: "Volgodonsk", Month: "january", DayNumber: "29", TempDay: "−6°", TempNight: "−7°", Condition: "Облачно"},
			w.DayWeather{City: "Volgodonsk", Month: "january", DayNumber: "30", TempDay: "−6°", TempNight: "−7°", Condition: "Облачно"},
			w.DayWeather{City: "Volgodonsk", Month: "january", DayNumber: "31", TempDay: "−5°", TempNight: "−7°", Condition: "Облачно"},
		},
	}
	encodedCSVData []byte = []byte(
		`days,Moscow,Volgodonsk
1 january,−2° −6° Облачно и слабый снег,−3° −6° Облачно
2 january,−5° −6° Облачно,−1° −3° Облачно
3 january,−7° −9° Облачно и слабый снег,−1° −4° Облачно и слабый снег
4 january,−8° −9° Облачно и слабый снег,−4° −6° Облачно
5 january,−8° −11° Облачно,−2° −5° Облачно
6 january,−10° −11° Облачно,−3° −4° Облачно и слабый снег
7 january,−10° −11° Облачно,−3° −6° Облачно и слабый снег
8 january,−9° −9° Облачно,−4° −5° Облачно
9 january,−8° −9° Облачно и слабый снег,−3° −5° Облачно
10 january,−7° −8° Облачно и слабый снег,−2° −3° Облачно
11 january,−6° −7° Облачно и слабый снег,−2° −4° Облачно
12 january,−5° −5° Облачно и слабый снег,−1° −2° Облачно
13 january,−3° −4° Облачно и слабый снег,+1° −1° Облачно
14 january,−3° −5° Облачно,0° −2° Облачно и слабый снег
15 january,−4° −6° Облачно и слабый снег,0° −1° Облачно
16 january,−5° −7° Облачно и слабый снег,0° −1° Облачно
17 january,−7° −9° Облачно,−1° −3° Облачно
18 january,−8° −10° Облачно и слабый снег,−2° −4° Облачно
19 january,−8° −10° Облачно,−2° −5° Облачно
20 january,−8° −10° Ясно,−4° −5° Облачно
21 january,−9° −11° Облачно,−3° −4° Облачно
22 january,−9° −9° Облачно и слабый снег,−2° −3° Облачно и слабый снег
23 january,−8° −9° Облачно,−1° −3° Ясно
24 january,−8° −10° Облачно,−3° −6° Облачно и слабый снег
25 january,−9° −12° Облачно,−5° −6° Облачно
26 january,−10° −11° Облачно и слабый снег,−5° −8° Облачно и слабый снег
27 january,−8° −9° Облачно,−7° −7° Облачно
28 january,−7° −9° Ясно,−5° −6° Облачно и слабый снег
29 january,−7° −9° Облачно и слабый снег,−6° −7° Облачно
30 january,−7° −8° Облачно,−6° −7° Облачно
31 january,−6° −8° Облачно и слабый снег,−5° −7° Облачно
`)
	encodedJSONData []byte = []byte(`[[{"City":"Moscow","Month":"january","DayNumber":"1","TempDay":"−2°","TempNight":"−6°","Condition":"Облачно и слабый снег"},{"City":"Moscow","Month":"january","DayNumber":"2","TempDay":"−5°","TempNight":"−6°","Condition":"Облачно"},{"City":"Moscow","Month":"january","DayNumber":"3","TempDay":"−7°","TempNight":"−9°","Condition":"Облачно и слабый снег"},{"City":"Moscow","Month":"january","DayNumber":"4","TempDay":"−8°","TempNight":"−9°","Condition":"Облачно и слабый снег"},{"City":"Moscow","Month":"january","DayNumber":"5","TempDay":"−8°","TempNight":"−11°","Condition":"Облачно"},{"City":"Moscow","Month":"january","DayNumber":"6","TempDay":"−10°","TempNight":"−11°","Condition":"Облачно"},{"City":"Moscow","Month":"january","DayNumber":"7","TempDay":"−10°","TempNight":"−11°","Condition":"Облачно"},{"City":"Moscow","Month":"january","DayNumber":"8","TempDay":"−9°","TempNight":"−9°","Condition":"Облачно"},{"City":"Moscow","Month":"january","DayNumber":"9","TempDay":"−8°","TempNight":"−9°","Condition":"Облачно и слабый снег"},{"City":"Moscow","Month":"january","DayNumber":"10","TempDay":"−7°","TempNight":"−8°","Condition":"Облачно и слабый снег"},{"City":"Moscow","Month":"january","DayNumber":"11","TempDay":"−6°","TempNight":"−7°","Condition":"Облачно и слабый снег"},{"City":"Moscow","Month":"january","DayNumber":"12","TempDay":"−5°","TempNight":"−5°","Condition":"Облачно и слабый снег"},{"City":"Moscow","Month":"january","DayNumber":"13","TempDay":"−3°","TempNight":"−4°","Condition":"Облачно и слабый снег"},{"City":"Moscow","Month":"january","DayNumber":"14","TempDay":"−3°","TempNight":"−5°","Condition":"Облачно"},{"City":"Moscow","Month":"january","DayNumber":"15","TempDay":"−4°","TempNight":"−6°","Condition":"Облачно и слабый снег"},{"City":"Moscow","Month":"january","DayNumber":"16","TempDay":"−5°","TempNight":"−7°","Condition":"Облачно и слабый снег"},{"City":"Moscow","Month":"january","DayNumber":"17","TempDay":"−7°","TempNight":"−9°","Condition":"Облачно"},{"City":"Moscow","Month":"january","DayNumber":"18","TempDay":"−8°","TempNight":"−10°","Condition":"Облачно и слабый снег"},{"City":"Moscow","Month":"january","DayNumber":"19","TempDay":"−8°","TempNight":"−10°","Condition":"Облачно"},{"City":"Moscow","Month":"january","DayNumber":"20","TempDay":"−8°","TempNight":"−10°","Condition":"Ясно"},{"City":"Moscow","Month":"january","DayNumber":"21","TempDay":"−9°","TempNight":"−11°","Condition":"Облачно"},{"City":"Moscow","Month":"january","DayNumber":"22","TempDay":"−9°","TempNight":"−9°","Condition":"Облачно и слабый снег"},{"City":"Moscow","Month":"january","DayNumber":"23","TempDay":"−8°","TempNight":"−9°","Condition":"Облачно"},{"City":"Moscow","Month":"january","DayNumber":"24","TempDay":"−8°","TempNight":"−10°","Condition":"Облачно"},{"City":"Moscow","Month":"january","DayNumber":"25","TempDay":"−9°","TempNight":"−12°","Condition":"Облачно"},{"City":"Moscow","Month":"january","DayNumber":"26","TempDay":"−10°","TempNight":"−11°","Condition":"Облачно и слабый снег"},{"City":"Moscow","Month":"january","DayNumber":"27","TempDay":"−8°","TempNight":"−9°","Condition":"Облачно"},{"City":"Moscow","Month":"january","DayNumber":"28","TempDay":"−7°","TempNight":"−9°","Condition":"Ясно"},{"City":"Moscow","Month":"january","DayNumber":"29","TempDay":"−7°","TempNight":"−9°","Condition":"Облачно и слабый снег"},{"City":"Moscow","Month":"january","DayNumber":"30","TempDay":"−7°","TempNight":"−8°","Condition":"Облачно"},{"City":"Moscow","Month":"january","DayNumber":"31","TempDay":"−6°","TempNight":"−8°","Condition":"Облачно и слабый снег"}],[{"City":"Volgodonsk","Month":"january","DayNumber":"1","TempDay":"−3°","TempNight":"−6°","Condition":"Облачно"},{"City":"Volgodonsk","Month":"january","DayNumber":"2","TempDay":"−1°","TempNight":"−3°","Condition":"Облачно"},{"City":"Volgodonsk","Month":"january","DayNumber":"3","TempDay":"−1°","TempNight":"−4°","Condition":"Облачно и слабый снег"},{"City":"Volgodonsk","Month":"january","DayNumber":"4","TempDay":"−4°","TempNight":"−6°","Condition":"Облачно"},{"City":"Volgodonsk","Month":"january","DayNumber":"5","TempDay":"−2°","TempNight":"−5°","Condition":"Облачно"},{"City":"Volgodonsk","Month":"january","DayNumber":"6","TempDay":"−3°","TempNight":"−4°","Condition":"Облачно и слабый снег"},{"City":"Volgodonsk","Month":"january","DayNumber":"7","TempDay":"−3°","TempNight":"−6°","Condition":"Облачно и слабый снег"},{"City":"Volgodonsk","Month":"january","DayNumber":"8","TempDay":"−4°","TempNight":"−5°","Condition":"Облачно"},{"City":"Volgodonsk","Month":"january","DayNumber":"9","TempDay":"−3°","TempNight":"−5°","Condition":"Облачно"},{"City":"Volgodonsk","Month":"january","DayNumber":"10","TempDay":"−2°","TempNight":"−3°","Condition":"Облачно"},{"City":"Volgodonsk","Month":"january","DayNumber":"11","TempDay":"−2°","TempNight":"−4°","Condition":"Облачно"},{"City":"Volgodonsk","Month":"january","DayNumber":"12","TempDay":"−1°","TempNight":"−2°","Condition":"Облачно"},{"City":"Volgodonsk","Month":"january","DayNumber":"13","TempDay":"+1°","TempNight":"−1°","Condition":"Облачно"},{"City":"Volgodonsk","Month":"january","DayNumber":"14","TempDay":"0°","TempNight":"−2°","Condition":"Облачно и слабый снег"},{"City":"Volgodonsk","Month":"january","DayNumber":"15","TempDay":"0°","TempNight":"−1°","Condition":"Облачно"},{"City":"Volgodonsk","Month":"january","DayNumber":"16","TempDay":"0°","TempNight":"−1°","Condition":"Облачно"},{"City":"Volgodonsk","Month":"january","DayNumber":"17","TempDay":"−1°","TempNight":"−3°","Condition":"Облачно"},{"City":"Volgodonsk","Month":"january","DayNumber":"18","TempDay":"−2°","TempNight":"−4°","Condition":"Облачно"},{"City":"Volgodonsk","Month":"january","DayNumber":"19","TempDay":"−2°","TempNight":"−5°","Condition":"Облачно"},{"City":"Volgodonsk","Month":"january","DayNumber":"20","TempDay":"−4°","TempNight":"−5°","Condition":"Облачно"},{"City":"Volgodonsk","Month":"january","DayNumber":"21","TempDay":"−3°","TempNight":"−4°","Condition":"Облачно"},{"City":"Volgodonsk","Month":"january","DayNumber":"22","TempDay":"−2°","TempNight":"−3°","Condition":"Облачно и слабый снег"},{"City":"Volgodonsk","Month":"january","DayNumber":"23","TempDay":"−1°","TempNight":"−3°","Condition":"Ясно"},{"City":"Volgodonsk","Month":"january","DayNumber":"24","TempDay":"−3°","TempNight":"−6°","Condition":"Облачно и слабый снег"},{"City":"Volgodonsk","Month":"january","DayNumber":"25","TempDay":"−5°","TempNight":"−6°","Condition":"Облачно"},{"City":"Volgodonsk","Month":"january","DayNumber":"26","TempDay":"−5°","TempNight":"−8°","Condition":"Облачно и слабый снег"},{"City":"Volgodonsk","Month":"january","DayNumber":"27","TempDay":"−7°","TempNight":"−7°","Condition":"Облачно"},{"City":"Volgodonsk","Month":"january","DayNumber":"28","TempDay":"−5°","TempNight":"−6°","Condition":"Облачно и слабый снег"},{"City":"Volgodonsk","Month":"january","DayNumber":"29","TempDay":"−6°","TempNight":"−7°","Condition":"Облачно"},{"City":"Volgodonsk","Month":"january","DayNumber":"30","TempDay":"−6°","TempNight":"−7°","Condition":"Облачно"},{"City":"Volgodonsk","Month":"january","DayNumber":"31","TempDay":"−5°","TempNight":"−7°","Condition":"Облачно"}]]`)
)
