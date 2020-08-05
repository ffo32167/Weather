package siteparse

import (
	"reflect"
	"testing"

	c "github.com/ffo32167/weather/weatherParser/config"
)

func TestChooseSiteParser(t *testing.T) {
	type args struct {
		site   string
		config *c.Config
	}
	tests := []struct {
		name       string
		args       args
		wantSource SiteParser
	}{
		{"basic worldWeather", args{"worldweather", &cfg}, worldWeather{cfg.WorldWeatherAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSource := ChooseSiteParser(tt.args.site, tt.args.config); !reflect.DeepEqual(gotSource, tt.wantSource) {
				t.Errorf("ChooseSiteParser() = %v, want %v", gotSource, tt.wantSource)
			}
		})
	}
}
