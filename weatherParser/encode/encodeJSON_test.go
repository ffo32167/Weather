package encode

import (
	"bytes"
	"reflect"
	"testing"

	w "github.com/ffo32167/weather/weatherParser/weatherdata"
)

func Test_json_Encode(t *testing.T) {
	type args struct {
		dataStruct [][]w.DayWeather
		cities     []string
	}
	tests := []struct {
		name            string
		j               json
		args            args
		wantEncodedData bytes.Buffer
		wantFormat      string
	}{
		{
			"csv_Encode",
			newJSON(),
			args{encodeData, []string{"Moscow", "Volgodonsk"}},
			*bytes.NewBuffer(encodedJSONData),
			".json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEncodedData, gotFormat := tt.j.Encode(tt.args.dataStruct, tt.args.cities)
			if !reflect.DeepEqual(gotEncodedData, tt.wantEncodedData) {
				t.Errorf("json.Encode() gotEncodedData = %v, want %v", gotEncodedData, tt.wantEncodedData)
				/*
					encodedDataFile, err := os.Create(`C:\encodedDataFile.json`)
					if err != nil {
						t.Error(`cant open C:\encodedDataFile.json : `, err)
					}
					defer encodedDataFile.Close()
					_, err = encodedDataFile.Write(gotEncodedData.Bytes())
					if err != nil {
						t.Error(`cant write gotEncodedData: `, err)
					}

					wantEncodedDataFile, err := os.Create(`C:\wantEncodedDataFile.json`)
					if err != nil {
						t.Error(`cant open C:\wantEncodedDataFile.json : `, err)
					}
					defer wantEncodedDataFile.Close()
					_, err = wantEncodedDataFile.Write(tt.wantEncodedData.Bytes())
					if err != nil {
						t.Error(`cant write gotEncodedData: `, err)
					}
				*/
			}
			if gotFormat != tt.wantFormat {
				t.Errorf("json.Encode() gotFormat = %v, want %v", gotFormat, tt.wantFormat)
			}
		})
	}
}
