package main

import (
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/matches/", GetUpcomingMatches)
	//mux.HandleFunc("/matches/live", GetLiveMatches)
	//mux.HandleFunc("/matches/top", GetTopTierMatches)
	//mux.HandleFunc("/matches/lan", GetLanMatches)

	http.ListenAndServe(":3000", mux)
}
