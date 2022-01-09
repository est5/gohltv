package main

import (
	mux "github.com/gorilla/mux"
	"net/http"
)

func (app *application) routes() http.Handler {
	r := mux.NewRouter().Methods("GET").Subrouter()

	r.HandleFunc("/matches/{type}", app.GetUpcomingMatches)
	r.HandleFunc("/matches", app.GetUpcomingMatches)
	r.HandleFunc("/news/{year}/{month}", app.GetNews)
	r.HandleFunc("/news", app.GetNews)
	r.HandleFunc("/results", app.GetResults)
	r.HandleFunc("/events/ongoing", app.GetOngoingEvents)
	r.HandleFunc("/events/upcoming", app.GetUpcomingEvents)
	//r.HandleFunc("/events/archive", app.GetFinishedEvents)
	//stats

	r.Use(app.loggingMiddleware, app.headersMiddleware)
	return r
}
