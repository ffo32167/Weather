package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	pb "github.com/ffo32167/weather/weatherProto"
	"github.com/sirupsen/logrus"
)

//	Загрузить параметры для приложения
func loadParams(appPath string) (params *pb.WeatherParams) {
	file, err := os.Open(filepath.Join(appPath, `params.json`))
	if err != nil {
		log.WithFields(logrus.Fields{"error": err}).Fatal("can't open params.json")
	}
	defer file.Close()
	buff, err := ioutil.ReadAll(file)
	if err != nil {
		log.WithFields(logrus.Fields{"error": err}).Fatal("can't read params.json")
	}
	err = json.Unmarshal(buff, &params)
	if err != nil {
		log.WithFields(logrus.Fields{"error": err}).Fatal("can't unmarshal params.json")
	}
	return
}

//	Развернуть срез месяцев []int{11,2}
//	в срез []string{"november", "december", "january", "february"}
func monthsParse(monthsNumbers []int32) (monthsNames []string) {
	calendarMonths := [12]string{"january", "february", "march", "april", "may", "june", "july", "august", "september", "october", "november", "december"}
	//	Проверить параметры
	if len(monthsNumbers) != 2 || monthsNumbers[0] < 1 || monthsNumbers[0] > 12 || monthsNumbers[1] < 1 || monthsNumbers[1] > 12 {
		log.WithFields(logrus.Fields{"monthsNumbers": monthsNumbers}).Fatal("incorrect interval of months")
	}
	//	если не сделать, то будет mismatched type int(Go) and int32(proto)  :(
	var i int32
	//	Если месяца по порядку, то вставить недостающее
	if monthsNumbers[1] > monthsNumbers[0] {
		for i = 0; i < monthsNumbers[1]-monthsNumbers[0]+1; i++ {
			monthsNames = append(monthsNames, calendarMonths[monthsNumbers[0]+i-1])
		}
		//	Если не по порядку, то вставить недостающие до конца года и с начала года
	} else if monthsNumbers[0] > monthsNumbers[1] {
		for i = monthsNumbers[0] - 1; i < 12; i++ {
			monthsNames = append(monthsNames, calendarMonths[i])
		}
		for i = 0; i < monthsNumbers[1]; i++ {
			monthsNames = append(monthsNames, calendarMonths[i])
		}
		//	Если месяц один, то его и вставить
	} else if monthsNumbers[0] == monthsNumbers[1] {
		monthsNames = append(monthsNames, calendarMonths[monthsNumbers[1]-1])
	}
	return monthsNames
}
