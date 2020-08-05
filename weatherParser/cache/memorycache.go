package cache

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	w "github.com/ffo32167/weather/weatherParser/weatherresponse"
	"github.com/sirupsen/logrus"
)

var errMemCacheMonthNotFound = errors.New("Month data not found in memory")

// Path Сформировать путь к кэшу
func Path(pathParts ...string) string {
	return filepath.Join(pathParts[0], `cache`, pathParts[1], pathParts[2], pathParts[3]+`_`+pathParts[4]+`_`+pathParts[5]+`.json`)
}

// Проверить наличие данных в кэше на диске
func cacheMonthCheck(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

// CacheMonthWrite сохраняет полученные и обработанные данные по месяцу на диск
func CacheMonthWrite(path string, data []w.WeatherResponse) {
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
func cacheOpen(path string) []w.WeatherResponse {
	file, err := os.Open(path)
	if err != nil {
		logrus.WithFields(logrus.Fields{"path": path, "err": err}).Error("can't open cache on path")
	}
	defer file.Close()
	buff, err := ioutil.ReadAll(file)
	if err != nil {
		logrus.WithFields(logrus.Fields{"path": path, "err": err}).Error("can't read cache on path")
	}
	var data []w.WeatherResponse
	err = json.Unmarshal(buff, &data)
	if err != nil {
		logrus.WithFields(logrus.Fields{"path": path, "err": err}).Error("can't unmarshal cache on path")
	}
	return data
}

// WeatherMemCache это кэш в памяти
type WeatherMemCache struct {
	mx    sync.RWMutex
	cache map[string][]w.WeatherResponse
}

// NewWeatherCache Инициализирует кэш в памяти
func NewWeatherCache() (wmc *WeatherMemCache) {
	wmc = &WeatherMemCache{
		cache: make(map[string][]w.WeatherResponse),
	}
	return wmc
}

// CacheRead получает данные месяца из памяти(city, month)
func (wmc *WeatherMemCache) CacheRead(path string) (wr []w.WeatherResponse, err error) {
	wmc.mx.RLock()
	defer wmc.mx.RUnlock()
	wr, ok := wmc.cache[path]
	if ok {
		logrus.Info("получены данные из кэша", path)
		return wr, nil
	}
	return nil, errMemCacheMonthNotFound
}

// CacheMonthWrite сохраняет данные за месяц в кэш и на диск
func (wmc *WeatherMemCache) CacheMonthWrite(path string, wr []w.WeatherResponse) {
	wmc.cacheMonthStore(path, wr)
	logrus.WithFields(logrus.Fields{"path": path, "weatherResponse len:": len(wr)}).Debug()
	CacheMonthWrite(path, wr)
}

// Сохранить данные месяца в память
func (wmc *WeatherMemCache) cacheMonthStore(path string, wr []w.WeatherResponse) {
	wmc.mx.Lock()
	defer wmc.mx.Unlock()
	logrus.WithFields(logrus.Fields{"path": path, "weatherResponse len:": len(wr)}).Debug()
	wmc.cache[path] = wr
}

// CacheLoad Читает папку с кэшем и загружает его в память
func (wmc *WeatherMemCache) CacheLoad(appPath string) {
	// создаём список файлов на загрузку
	fileList := make([]string, 0)
	filesPath := filepath.Join(appPath, "cache")
	filepath.Walk(filesPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(info.Name()) == ".json" {
			fileList = append(fileList, path)
		}
		return nil
	})
	// загружаем список в память
	var wg sync.WaitGroup
	wg.Add(len(fileList))
	for _, file := range fileList {
		go func(file string) {
			defer wg.Done()
			fileCache := cacheOpen(file)
			wmc.cacheMonthStore(file, fileCache)
		}(file)
	}
	wg.Wait()
	logrus.WithFields(logrus.Fields{"cache loaded len": len(wmc.cache)}).Info()
}
