package main

import "flag"

// Структура для промежуточного хранения данных
type weatherResponse struct {
	City      string
	Month     string
	DayNumber string
	TempDay   string
	TempNight string
	Condition string
}

/* todo:
		улучшить работу с ошибками: в методе cacheLoad при ошибке чтения files, err := ioutil.ReadDir(path)
				создать специальный тип ошибки(путь не найден), и если именно эта ошибка происходит, то
				создавать дерево папок аналогично функции cacheMonthWrite
		реализовать shutdown и интерцепторы для grpc и http
		разбить на пакеты?
-----------------------------------------------------------------------------------------------------------
		добавить модуль сборник типовых действий не вошедших сюда - работа с каналами, контекстом и т.д.
		80% code test coverage
*/

func main() {
	var mode string
	flag.StringVar(&mode, "mode", "", "mode: s - grpc service; application is a default mode")
	flag.Parse()
	config := newConfig()
	switch mode {
	case "s":
		runService(config)
	default:
		runApp(config)
	}
}
