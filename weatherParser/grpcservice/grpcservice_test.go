package grpcservice

import (
	"reflect"
	"testing"

	ch "github.com/ffo32167/weather/weatherParser/cache"
	c "github.com/ffo32167/weather/weatherParser/config"
	s "github.com/ffo32167/weather/weatherParser/siteparse"
	w "github.com/ffo32167/weather/weatherParser/weatherresponse"
	pb "github.com/ffo32167/weather/weatherProto"
)

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

func TestWeatherDataGrab(t *testing.T) {
	type args struct {
		sp     s.SiteParser
		params *pb.WeatherParams
		config *c.Config
		cr     ch.CacheReader
	}
	tests := []struct {
		name   string
		args   args
		wantWr [][]w.WeatherResponse
	}{
		{
			"WeatherDataGrab",
			args{
				s.ChooseSiteParser("yandex", &cfg),
				param,
				&cfg,
				memoryCache,
			},
			weatherDataGrabResult,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotWr := WeatherDataGrab(tt.args.sp, tt.args.params, tt.args.config, tt.args.cr); !reflect.DeepEqual(gotWr, tt.wantWr) {
				t.Errorf("WeatherDataGrab() = %v, want %v", gotWr, tt.wantWr)
			}
		})
	}
}
