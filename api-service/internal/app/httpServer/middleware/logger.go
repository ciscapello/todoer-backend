package middleware

import (
	"github.com/charmbracelet/log"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info(r.Method, "URL", r.URL)
		next.ServeHTTP(w, r)
	})
}
