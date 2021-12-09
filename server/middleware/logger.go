package middleware

import (
	"fmt"
	"log"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		printHTTPLog(w, r)
	})
}

func printHTTPLog(w http.ResponseWriter, r *http.Request) {
	var bytesStr string
	contentLength := w.Header().Get("Content-Length")
	if contentLength != "" {
		bytesStr = fmt.Sprintf("(%s bytes)", contentLength)
	}

	var separator string
	statusCode := w.Header().Get("X-Status-Code")
	if statusCode != "" || bytesStr != "" {
		separator = "->"
	}

	log.Println(r.Method, r.URL.Path, separator, statusCode, bytesStr)
}
