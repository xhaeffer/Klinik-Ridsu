package session

import (
	"encoding/gob"
	"net/http"

	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

func init() {
	secretKey := []byte("klinikridsu")
	store = sessions.NewCookieStore(secretKey)
	store.Options.HttpOnly = false
	gob.Register(map[string]interface{}{})
}

func SetSession(w http.ResponseWriter, r *http.Request, key string, userData map[string]interface{}) error {
	session, err := store.Get(r, "user")
	if err != nil {
		return err
	}

	if _, exists := session.Values[key]; !exists {
		session.Values[key] = make(map[string]interface{})
	}

	userSession := session.Values[key].(map[string]interface{})
	for key, value := range userData {
		userSession[key] = value
	}

	session.Options.MaxAge = 3600
	
	err = session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}

func GetSession(r *http.Request) map[interface{}]interface{} {
	session, _ := store.Get(r, "user")
	return session.Values
}

func ClearSession(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	session.Options.MaxAge = -1
	session.Save(r, w)
}
