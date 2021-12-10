package server

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	pkgPage "github.com/michaeltelford/swoop/page"
	pkgMiddleware "github.com/michaeltelford/swoop/server/middleware"
)

func NewServer(address string, pages []pkgPage.IPage) *http.Server {
	mux := NewMux(pages)

	return &http.Server{
		Addr:         address,
		Handler:      buildMiddleware(mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
}

func NewMux(pages []pkgPage.IPage) *http.ServeMux {
	mux := http.NewServeMux()
	registerPages(mux, pages)

	return mux
}

func registerPages(mux *http.ServeMux, pages []pkgPage.IPage) {
	for _, page := range pages {
		registerPage(mux, page)
	}
}

func registerPage(mux *http.ServeMux, page pkgPage.IPage) {
	mux.HandleFunc(page.Route(), func(w http.ResponseWriter, r *http.Request) {
		respondWithContent(w, page.Content())
	})
}

func respondWithContent(w http.ResponseWriter, content template.HTML) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(content)))
	w.Header().Set("X-Status-Code", "200")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(content))
}

func buildMiddleware(mux *http.ServeMux) http.Handler {
	return pkgMiddleware.Logger(pkgMiddleware.HTTPMethod(mux))
}
