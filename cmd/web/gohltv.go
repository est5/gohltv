package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

type application struct {
	log *log.Logger
}

func main() {
	mux := http.NewServeMux()
	app := application{log: log.New()}
	mux.HandleFunc("/matches/", app.GetUpcomingMatches)
	//mux.HandleFunc("/matches/live", GetLiveMatches)
	//mux.HandleFunc("/matches/top", GetTopTierMatches)
	//mux.HandleFunc("/matches/lan", GetLanMatches)

	log.Infoln("Starting at port :3000")
	http.ListenAndServe(":3000", mux)
}
