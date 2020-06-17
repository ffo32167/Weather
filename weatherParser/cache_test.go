package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var appPath = ``

func TestGetCachePath(t *testing.T) {
	type testCase struct {
		AppPath string
		Source  string `json:"source"`
		Country string `json:"country"`
		City    string `json:"city"`
		Month   string `json:"month"`
		Year    string `json:"year"`
		Result  string `json:"result"`
	}
	testCases := make([]testCase, 0)
	file, err := os.Open(filepath.Join(appPath, `testData`, `TestGetCachePath.json`))
	if err != nil {
		t.Errorf("can't open getCachePath.json")
	}
	defer file.Close()
	buff, err := ioutil.ReadAll(file)
	if err != nil {
		t.Errorf("can't read getCachePath.json")
	}
	err = json.Unmarshal(buff, &testCases)
	if err != nil {
		t.Errorf("can't unmarshal getCachePath.json")
	}
	var testResult, funcResultCachePath string

	for _, val := range testCases {
		funcResultCachePath = cachePath(appPath, val.Source, val.Country, val.City, val.Month, val.Year)
		testResult = appPath + val.Result
		fmt.Println(val.Result)
		if testResult != funcResultCachePath {
			t.Errorf("TestGetCachePath: expect \n %v \n have \n %v \n", testResult, funcResultCachePath) //
		}
	}
}

func TestCheckCacheData(t *testing.T) {
	type testCase struct {
		path       string
		testResult bool
	}
	testCases := []testCase{
		{filepath.Join(appPath, `cache`, `yandex`, `russia`, `moscow_february_2018.json`), true},
		{filepath.Join(appPath, `cache`, `yandex`, `russia`, `moscow_november_2018.json`), true},
	}
	for _, val := range testCases {
		funcResult := cacheMonthCheck(val.path)
		if funcResult != val.testResult {
			t.Errorf("TestCheckCacheData: expect %v have %v", val.testResult, funcResult)
		}
	}
}

func TestGetDataCache(t *testing.T) {
	type testCase struct {
		path       string
		testResult []weatherResponse
	}
	testCases := []testCase{
		{path: filepath.Join(appPath, `testData`, `moscow_february_2018.json`),
			testResult: []weatherResponse{
				{City: "moscow", Month: "february", DayNumber: "1", TempDay: "−6°", TempNight: "−7°", Condition: "Облачно и слабый снег"},
				{City: "moscow", Month: "february", DayNumber: "2", TempDay: "−5°", TempNight: "−6°", Condition: "Облачно и слабый снег"}},
		},
	}
	for _, val := range testCases {
		funcResult := cacheOpen(val.path)
		for i := range funcResult {
			if funcResult[i] != val.testResult[i] {
				t.Errorf("TestGetDataCache: expect %v have %v on element %v", val.testResult, funcResult, i)
			}
		}
	}
}

func TestCacheData(t *testing.T) {
	// Если мы кэшируем данные, то мы должны уметь их "раскэшировать"
	type testCase struct {
		path    string
		country string
		source  string
		city    string
		month   string
		year    string
		data    []weatherResponse
	}
	testCases := []testCase{
		{
			filepath.Join(appPath, `testData`),
			"russia",
			"yandex",
			"moscow",
			"february",
			"2018",
			[]weatherResponse{
				{City: "moscow", Month: "february", DayNumber: "1", TempDay: "−6°", TempNight: "−7°", Condition: "Облачно и слабый снег"},
				{City: "moscow", Month: "february", DayNumber: "2", TempDay: "−5°", TempNight: "−6°", Condition: "Облачно и слабый снег"}},
		},
	}
	for i, val := range testCases {
		err := os.RemoveAll(filepath.Join(appPath, `testData`, `cache`))
		if err != nil {
			t.Error(`can't clear testData\cache:`, err)
		}
		cachePath := cachePath(val.path, val.source, val.country, val.city, val.month, val.year)
		cacheMonthWrite(cachePath, val.data)
		if !cacheMonthCheck(cachePath) {
			t.Errorf("TestCacheData: can't find test result on element %v", i)
		}
		funcResult := cacheOpen(cachePath)
		for i := range funcResult {
			if funcResult[i] != val.data[i] {
				t.Errorf("TestCacheData: expect %v have %v on element %v", val.data, funcResult, i)
			}
		}
	}
}
