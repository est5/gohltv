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
	Lan       bool
	MatchTime time.Time
}

func GetUpcomingMatches(w http.ResponseWriter, r *http.Request) {
	var matches []Match
	c := colly.NewCollector()
	c.OnHTML("div.upcomingMatchesContainer", func(e *colly.HTMLElement) {
		e.ForEach("div.upcomingMatch", func(i int, element *colly.HTMLElement) {
			link := "https://www.hltv.org" + element.ChildAttr("a.match", "href")
			stars := element.Attr("stars")
			team1 := element.ChildText("div.matchTeam.team1") + " " + element.Attr("team1")
			team2 := element.ChildText("div.matchTeam.team2") + " " + element.Attr("team2")
			lan, _ := strconv.ParseBool(element.Attr("lan"))
			matchTime, _ := strconv.ParseInt(strings.TrimSpace(element.ChildAttr("div.matchTime", "data-unix")), 10, 64)
			date := time.UnixMilli(matchTime).UTC()

			m := Match{Link: link, Stars: stars, Team1: team1, Team2: team2, Lan: lan, MatchTime: date}
			matches = append(matches, m)
		})
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

//func GetLiveMatches(w http.ResponseWriter, r *http.Request) {
//	var matches []Match
//	c := colly.NewCollector()
//	c.OnHTML("div.liveMatchesContainer", func(e *colly.HTMLElement) {
//		e.ForEach("div.liveMatch-container", func(i int, element *colly.HTMLElement) {
//			link := "https://www.hltv.org" + element.ChildAttr("a.match", "href")
//			stars := element.Attr("stars")
//			team1 := element.ChildText("div.matchTeamName") + " " + element.Attr("team1")
//			team2 := element.ChildText("div.matchTeamName") + " " + element.Attr("team2")
//			lan, _ := strconv.ParseBool(element.Attr("lan"))
//
//			m := Match{Link: link, Stars: stars, Team1: team1, Team2: team2, Lan: lan}
//			matches = append(matches, m)
//		})
//	})
//	c.Visit("https://www.hltv.org/matches")
//
//	js, err := json.Marshal(matches)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	w.Write(js)
//}
