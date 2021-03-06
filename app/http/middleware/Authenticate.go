package middleware

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/sarulabs/di/v2"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			store := di.Get(r, "session").(*sessions.CookieStore)
			sess, err := store.Get(r, "gostalt")

			user := sess.Values["user"]
			if user == "" || user == nil || err != nil {
				http.Redirect(w, r, "/login", 302)
				return
			}

			next.ServeHTTP(w, r)
		},
	)
}
