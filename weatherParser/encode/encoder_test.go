package encode

import (
	"reflect"
	"testing"
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
