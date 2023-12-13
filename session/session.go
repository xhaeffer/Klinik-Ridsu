package session

import (
	"net/http"
	"fmt"

	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

func init() {
	secretKey := []byte("klinikridsu")
	store = sessions.NewCookieStore(secretKey)
}

// SetSession mengatur nilai pada sesi
func SetSession(w http.ResponseWriter, r *http.Request, key string, value interface{}) error {
	session, err := store.Get(r, "sesi_pengguna")
	if err != nil {
		return err
	}

	session.Values[key] = value
	session.Save(r, w)
	fmt.Println("Session values:", session.Values)


	return nil
}

// GetSession mendapatkan nilai dari sesi
func GetSession(r *http.Request, key string) interface{} {
	session, _ := store.Get(r, "sesi_pengguna")
	return session.Values[key]
}

// ClearSession membersihkan sesi
func ClearSession(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "sesi_pengguna")
	session.Options.MaxAge = -1
	session.Save(r, w)
}
