package main

import (
	mux "github.com/gorilla/mux"
	"net/http"
)

func (app *application) routes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/matches/{type}", app.GetUpcomingMatches).Methods("GET")
	r.HandleFunc("/matches", app.GetUpcomingMatches).Methods("GET")
	return r
}
