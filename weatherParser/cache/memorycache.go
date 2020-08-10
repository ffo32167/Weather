package cache

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	w "github.com/ffo32167/weather/weatherParser/weatherdata"
	"github.com/sirupsen/logrus"
)

// WeatherMemCache это кэш в памяти
type WeatherMemCache struct {
	mx    sync.RWMutex
	cache map[string][]w.DayWeather
}

// NewWeatherMemCache Инициализирует кэш в памяти
func NewWeatherMemCache() (wmc *WeatherMemCache) {
	wmc = &WeatherMemCache{
		cache: make(map[string][]w.DayWeather),
	}
	return wmc
}

// Создаёт список файлов на загрузку
func makeFilelist(path string) []string {
	fileList := make([]string, 0)
	filesPath := filepath.Join(path, "cache")
	filepath.Walk(filesPath, func(fpath string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(info.Name()) == ".json" {
			fileList = append(fileList, fpath)
		}
		return nil
	})
	return fileList
}

// Load WeatherMemCache: Читает папку с кэшем и загружает его в память
func (wmc *WeatherMemCache) Load(path string) {
	fileList := makeFilelist(path)
	// загружаем список в память
	var wg sync.WaitGroup
	wg.Add(len(fileList))
	for _, file := range fileList {
		go func(file string) {
			defer wg.Done()
			fileCache := cacheOpen(file)
			wmc.monthStore(file, fileCache)
		}(file)
	}
	wg.Wait()
	logrus.WithFields(logrus.Fields{"cache loaded len": len(wmc.cache)}).Info()
}

// Path создаёт путь к данным в кэше
func (wmc *WeatherMemCache) Path(pathParts ...string) string {
	return path(pathParts...)
}

// MonthRead получает данные месяца из памяти(city, month)
func (wmc *WeatherMemCache) MonthRead(path string) (wr []w.DayWeather, err error) {
	wmc.mx.RLock()
	defer wmc.mx.RUnlock()
	wr, ok := wmc.cache[path]
	if ok {
		logrus.Info("получены данные из кэша", path)
		return wr, nil
	}
	return nil, errMemCacheMonthNotFound
}

// MonthWrite сохраняет данные за месяц в кэш и на диск
func (wmc *WeatherMemCache) MonthWrite(path string, wr []w.DayWeather) {
	wmc.monthStore(path, wr)
	logrus.WithFields(logrus.Fields{"path": path, "DayWeather len:": len(wr)}).Debug()
	monthWrite(path, wr)
}

// Сохранить данные месяца в память
func (wmc *WeatherMemCache) monthStore(path string, wr []w.DayWeather) {
	wmc.mx.Lock()
	defer wmc.mx.Unlock()
	logrus.WithFields(logrus.Fields{"path": path, "DayWeather len:": len(wr)}).Debug()
	wmc.cache[path] = wr
}

// Cохранить полученные и обработанные данные по месяцу на диск
func monthWrite(path string, data []w.DayWeather) {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, 0)
	if err != nil {
		logrus.WithFields(logrus.Fields{"dir": dir, "err": err}).Error("can't create cache directory")
	}
	file, err := os.Create(path)
	if err != nil {
		logrus.WithFields(logrus.Fields{"path": path, "err": err}).Error("can't create cache file")
	}
	defer file.Close()
	buff, err := json.Marshal(data)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("can't marshal cache, error")
	}
	_, err = file.Write(buff)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("can't write cache data to file")
	}
}

// Извлечь данные по месяцу из кэша на диске
func cacheOpen(path string) []w.DayWeather {
	file, err := os.Open(path)
	if err != nil {
		logrus.WithFields(logrus.Fields{"path": path, "err": err}).Error("can't open cache on path")
	}
	defer file.Close()
	buff, err := ioutil.ReadAll(file)
	if err != nil {
		logrus.WithFields(logrus.Fields{"path": path, "err": err}).Error("can't read cache on path")
	}
	var data []w.DayWeather
	err = json.Unmarshal(buff, &data)
	if err != nil {
		logrus.WithFields(logrus.Fields{"path": path, "err": err}).Error("can't unmarshal cache on path")
	}
	return data
}

// Path Сформировать путь к кэшу
func path(pathParts ...string) string {
	return filepath.Join(pathParts[0], `cache`, pathParts[1], pathParts[2], pathParts[3]+`_`+pathParts[4]+`_`+pathParts[5]+`.json`)
}

var errMemCacheMonthNotFound = errors.New("Month data not found in memory")
