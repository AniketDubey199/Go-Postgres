package middleware

import (
	"fmt"
	"net/http"
)

func LoggingDetails(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Print(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
