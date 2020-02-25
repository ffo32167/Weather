package main

import (
	"math/rand"
	"net/http"
	"sync"
	// log "github.com/sirupsen/logrus"
)

//	Сессия внутренняя
type session struct {
	Login string
}

//	Сессия в cookies
type sessionID struct {
	ID string
}

//	Длина сессии в cookie
const sessKeyLen = 10

//	HTTP сервер асинхронный, а мапки нет, поэтому писать и читать список сессий
//	только через mutex читающепишущий
type sessionManager struct {
	mu       sync.RWMutex
	sessions map[sessionID]*session
}

var sessManager *sessionManager

//	Новый менеджер сессий, с новым mutex'ом
func newSessionManager() *sessionManager {
	return &sessionManager{
		mu:       sync.RWMutex{},
		sessions: map[sessionID]*session{},
	}
}

//	Через mutex создать новую сессию с помощью генерации случайной строки
func (sm *sessionManager) Create(in *session) (*sessionID, error) {
	sm.mu.Lock()
	id := sessionID{randStringRunes(sessKeyLen)}
	sm.mu.Unlock()
	sm.sessions[id] = in
	return &id, nil
}

// Проверить cookies на session
func checkSession(r *http.Request) (*session, error) {
	//	Взять cookie на пробу у пришедшего
	cookieSessionID, err := r.Cookie("session_id")
	//	Если ошибка "cookie нет"
	if err == http.ErrNoCookie {
		return nil, nil
		//	Если ошибка другая
	} else if err != nil {
		return nil, err
	}
	//	Проверить, что пришедшая cookie есть в нашем списке
	sess := sessManager.Check(&sessionID{
		ID: cookieSessionID.Value,
	})
	return sess, nil
}

//	Проверить, что cookie есть в списке
func (sm *sessionManager) Check(in *sessionID) *session {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	if sess, ok := sm.sessions[*in]; ok {
		return sess
	}
	return nil
}

//	В случае выхода удалить cookie
func (sm *sessionManager) Delete(in *sessionID) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.sessions, *in)
}

//	Список доступных символов для cookie
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//	Генерировать cookie длинны n
func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
