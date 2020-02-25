package main

import (
	"path/filepath"
	"testing"
)

func TestEncode(t *testing.T) {
	type testDataStruct struct {
		// путь к файлу с тестовыми данными
		testDataFilePath []string
		// список городов для теста
		testCities []string
		// позиции отдельных символов в результате для сравнения с результатами теста
		testResult map[int]int
		// размер результата теста
		testResultLenght int
		// результат тестирования функции
		funcResult []byte
	}
	var (
		cityWeather []weatherResponse
		wResponse   [][]weatherResponse
	)
	testData := []testDataStruct{
		{[]string{
			filepath.Join(appPath, `testdata`, `moscow_february_2018.json`),
			filepath.Join(appPath, `testdata`, `volgodonsk_february_2018.json`),
		},
			[]string{`moscow`, `volgodonsk`},
			make(map[int]int),
			0,
			make([]byte, 0),
		},
	}
	// по списку тестовых случаев
	for _, valTestData := range testData {
		// в каждом случае читаем информацию из файлов с тестовыми данными
		for _, valDataFilePath := range valTestData.testDataFilePath {
			cityWeather = cacheOpen(valDataFilePath)
			wResponse = append(wResponse, cityWeather)
		}
		funcResult := encode(wResponse, valTestData.testCities)
		valTestData.funcResult = funcResult.Bytes()

	}
}
