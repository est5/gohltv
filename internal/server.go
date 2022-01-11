package internal

import (
	"flag"
	mux "github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type application struct {
	log *log.Logger
}

func Run() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	app := application{log: log.New()}

	srv := &http.Server{
		Addr:         *addr,
		Handler:      app.routes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Infof("Starting serve on %s\n", *addr)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func (app *application) routes() http.Handler {
	r := mux.NewRouter().Methods("GET").Subrouter()

	r.HandleFunc("/matches/live", app.GetLiveMatches)
	r.HandleFunc("/matches/{type}", app.GetUpcomingMatches)
	r.HandleFunc("/matches", app.GetUpcomingMatches)

	r.HandleFunc("/news/{year}/{month}", app.GetNews)
	r.HandleFunc("/news", app.GetNews)

	r.HandleFunc("/results", app.GetResults)

	r.HandleFunc("/events/ongoing", app.GetOngoingEvents)
	r.HandleFunc("/events/upcoming", app.GetUpcomingEvents)
	r.HandleFunc("/events/archive", app.GetArchiveEvents)

	r.HandleFunc("/stats/players", app.GetStatsPlayers)
	r.HandleFunc("/stats/players/flashbangs", app.GetStatsPlayersFlashes)
	r.HandleFunc("/stats/players/openingkills", app.GetStatsPlayersOpeners)
	r.HandleFunc("/stats/players/pistols", app.GetStatsPlayersPistols)

	//result for particular event and match
	//forum?

	r.Use(app.loggingMiddleware, app.headersMiddleware)
	return r
}
