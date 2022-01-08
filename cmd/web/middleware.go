package main

import "net/http"

func (app *application) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.log.Infof("Request to %v", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func (app *application) headersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("User-Agent", RandomString())
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "deny")
		next.ServeHTTP(w, r)
	})
}
