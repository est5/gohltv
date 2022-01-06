package main

import (
	"encoding/json"
	"github.com/gocolly/colly"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Match struct {
	Link      string
	Stars     string
	Team1     string
	Team2     string
	MatchTime time.Time
}

func GetUpcomingMatches(w http.ResponseWriter, r *http.Request) {
	c := colly.NewCollector()
	var matches []Match
	c.OnHTML("div.upcomingMatchesContainer", func(e *colly.HTMLElement) {
		e.ForEach("div.upcomingMatch", func(i int, element *colly.HTMLElement) {
			link := "https://www.hltv.org" + element.ChildAttr("a.match", "href")
			stars := element.Attr("stars")
			team1 := element.ChildText("div.matchTeam.team1") + " " + element.Attr("team1")
			team2 := element.ChildText("div.matchTeam.team2") + " " + element.Attr("team2")
			matchTime, _ := strconv.ParseInt(strings.TrimSpace(element.ChildAttr("div.matchTime", "data-unix")), 10, 64)
			date := time.UnixMilli(matchTime).UTC()

			m := Match{Link: link, Stars: stars, Team1: team1, Team2: team2, MatchTime: date}
			matches = append(matches, m)
		})
	})
	c.OnRequest(func(request *colly.Request) {
		request.Headers.Set("User-Agent", RandomString())
	})

	c.Visit("https://www.hltv.org/matches")
	js, err := json.Marshal(matches)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
