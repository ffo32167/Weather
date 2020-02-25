package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/sirupsen/logrus"
)

// Сформировать путь к кэшу
func cachePath(pathParts ...string) string {
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

// Сохранить полученные и обработанные данные по месяцу на диск
func cacheMonthWrite(path string, data []weatherResponse) {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, 0)
	if err != nil {
		log.WithFields(logrus.Fields{"dir": dir, "err": err}).Error("can't create cache directory")
	}
	file, err := os.Create(path)
	if err != nil {
		log.WithFields(logrus.Fields{"path": path, "err": err}).Error("can't create cache file")
	}
	defer file.Close()
	buff, err := json.Marshal(data)
	if err != nil {
		log.WithFields(logrus.Fields{"err": err}).Error("can't marshal cache, error")
	}
	_, err = file.Write(buff)
	if err != nil {
		log.WithFields(logrus.Fields{"err": err}).Error("can't write cache data to file")
	}
}

// Извлечь данные по месяцу из кэша на диске
func cacheOpen(path string) []weatherResponse {
	file, err := os.Open(path)
	if err != nil {
		log.WithFields(logrus.Fields{"path": path, "err": err}).Error("can't open cache on path")
	}
	defer file.Close()
	buff, err := ioutil.ReadAll(file)
	if err != nil {
		log.WithFields(logrus.Fields{"path": path, "err": err}).Error("can't read cache on path")
	}
	var data []weatherResponse
	err = json.Unmarshal(buff, &data)
	if err != nil {
		log.WithFields(logrus.Fields{"path": path, "err": err}).Error("can't unmarshal cache on path")
	}
	return data
}

// Извлечь данные по месяцу из кэша на диске
func (app) cacheRead(path string) []weatherResponse {
	return cacheOpen(path)
}

// Сохранить полученные и обработанные данные по месяцу на диск
func (app) cacheMonthWrite(path string, data []weatherResponse) {
	cacheMonthWrite(path, data)
}

// Кэш в памяти
type weatherMemCache struct {
	mx    sync.RWMutex
	cache map[string][]weatherResponse
}

// Инициализировать кэш в памяти
func newWeatherCache() weatherMemCache {
	return weatherMemCache{
		cache: make(map[string][]weatherResponse),
		// Мьютекс никак инициализировать не нужно, его "нулевое значение" это разлоченный мьютекс,
		// готовый к использованию
	}
}

// Получить данные месяца из памяти(city, month)
func (wmc *weatherMemCache) cacheRead(path string) []weatherResponse {
	wmc.mx.RLock()
	defer wmc.mx.RUnlock()
	log.WithFields(logrus.Fields{"path": path}).Debug()
	wr, _ := wmc.cache[path]
	return wr
}

func (wmc *weatherMemCache) cacheMonthWrite(path string, wr []weatherResponse) {
	wmc.cacheMonthStore(path, wr)
	log.WithFields(logrus.Fields{"path": path, "weatherResponse len:": len(wr)}).Debug()
	cacheMonthWrite(path, wr)
}

// Сохранить данные месяца в память
func (wmc *weatherMemCache) cacheMonthStore(path string, wr []weatherResponse) {
	wmc.mx.Lock()
	defer wmc.mx.Unlock()
	log.WithFields(logrus.Fields{"path": path, "weatherResponse len:": len(wr)}).Debug()
	wmc.cache[path] = wr
}

// Прочитать папку с кэшем и загрузить его в память
func (wmc *weatherMemCache) cacheLoad(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	var (
		wg       sync.WaitGroup
		fileName string
	)
	wg.Add(len(files))
	for _, file := range files {
		fileName = file.Name()
		// Взять название города и месяца из названия файла
		// Прочитать файл и загрузить данные в память
		go func(fileName string) {
			defer wg.Done()
			fullPath := filepath.Join(path, fileName)
			fileCache := cacheOpen(fullPath)
			wmc.cacheMonthStore(fullPath, fileCache)
		}(fileName)
	}
	wg.Wait()
	log.WithFields(logrus.Fields{"cache loaded len": len(wmc.cache)}).Info()
}
