package main

import (
	mux "github.com/gorilla/mux"
	"net/http"
)

func (app *application) routes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/matches/{type}", app.GetUpcomingMatches).Methods("GET")
	r.HandleFunc("/matches", app.GetUpcomingMatches).Methods("GET")
	r.HandleFunc("/news/{year}/{month}", app.GetNews).Methods("GET")
	r.HandleFunc("/news", app.GetNews).Methods("GET")
	r.HandleFunc("/results", app.GetResults).Methods("GET")
	r.HandleFunc("/events/ongoing", app.GetOngoingEvents).Methods("GET")
	r.HandleFunc("/events/upcoming", app.GetUpcomingEvents).Methods("GET")
	//r.HandleFunc("/events/archive", app.GetFinishedEvents).Methods("GET")
	r.Use(app.loggingMiddleware, app.headersMiddleware)
	return r
}
