package swoop

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	pkgPages "github.com/michaeltelford/swoop/pages"
)

func NewServer(address string, pages []pkgPages.IPage) *http.Server {
	mux := NewMux(pages)

	return &http.Server{
		Addr:         address,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
}

func NewMux(pages []pkgPages.IPage) *http.ServeMux {
	mux := http.NewServeMux()
	registerPages(mux, pages)

	return mux
}

func registerPages(mux *http.ServeMux, pages []pkgPages.IPage) {
	for _, page := range pages {
		registerPage(mux, page)
	}
}

func registerPage(mux *http.ServeMux, page pkgPages.IPage) {
	mux.HandleFunc(page.Route(), func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			respondWithPageContent(w, page)
		} else {
			respondWithStatusNotAllowed(w)
		}
	})
}

func respondWithPageContent(w http.ResponseWriter, page pkgPages.IPage) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(page.Content())))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(page.Content()))
}

func respondWithStatusNotAllowed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte(fmt.Sprintf("%d method not allowed", http.StatusMethodNotAllowed)))
}
