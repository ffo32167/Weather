package encode

import (
	"bytes"
	"reflect"
	"testing"

	w "github.com/ffo32167/weather/weatherParser/weatherdata"
)

func Test_csv_Encode(t *testing.T) {
	type args struct {
		data   [][]w.DayWeather
		cities []string
	}
	tests := []struct {
		name            string
		c               csv
		args            args
		wantEncodedData bytes.Buffer
		wantFormat      string
	}{
		{
			"csv_Encode",
			newCSV(),
			args{encodeData, []string{"Moscow", "Volgodonsk"}},
			*bytes.NewBuffer(encodedCSVData),
			".csv",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEncodedData, gotFormat := tt.c.Encode(tt.args.data, tt.args.cities)
			if !reflect.DeepEqual(gotEncodedData, tt.wantEncodedData) {
				t.Errorf("csv.Encode() gotEncodedData = %v, want %v", gotEncodedData, tt.wantEncodedData)
				/*
					encodedDataFile, err := os.Create(`C:\encodedDataFile.txt`)
					if err != nil {
						t.Error(`cant open C:\encodedDataFile.txt : `, err)
					}
					defer encodedDataFile.Close()
					_, err = encodedDataFile.Write(gotEncodedData.Bytes())
					if err != nil {
						t.Error(`cant write gotEncodedData: `, err)
					}

					wantEncodedDataFile, err := os.Create(`C:\wantEncodedDataFile.txt`)
					if err != nil {
						t.Error(`cant open C:\wantEncodedDataFile.txt : `, err)
					}
					defer wantEncodedDataFile.Close()
					_, err = wantEncodedDataFile.Write(tt.wantEncodedData.Bytes())
					if err != nil {
						t.Error(`cant write gotEncodedData: `, err)
					}
				*/
			}
			if gotFormat != tt.wantFormat {
				t.Errorf("csv.Encode() gotFormat = %v, want %v", gotFormat, tt.wantFormat)
			}
		})
	}
}
