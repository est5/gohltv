package main

import "net/http"

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/matches/upcoming", GetUpcomingMatches)
	//mux.HandleFunc("/matches/live", GetLiveMatches)

	http.ListenAndServe(":3000", mux)
}
