package middleware

import (
	"fmt"
	"net/http"
)

func HTTPMethod(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			respondWithMethodNotAllowed(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func respondWithMethodNotAllowed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte(fmt.Sprintf("%d method not allowed", http.StatusMethodNotAllowed)))
}
