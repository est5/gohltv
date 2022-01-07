package main

import (
	"encoding/json"
	"github.com/est5/gohltv/pkg/models"
	"github.com/gocolly/colly"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (app *application) GetUpcomingMatches(w http.ResponseWriter, r *http.Request) {
	c := colly.NewCollector()
	vars := mux.Vars(r)
	url := getLink(vars["type"])
	var matches []models.UpcomingMatch

	c.OnHTML("div.upcomingMatchesContainer", func(e *colly.HTMLElement) {
		e.ForEach("div.upcomingMatch", func(i int, element *colly.HTMLElement) {
			link := "https://www.hltv.org" + element.ChildAttr("a.match", "href")
			stars := element.Attr("stars")
			team1 := element.ChildText("div.matchTeam.team1") + " " + element.Attr("team1")
			team2 := element.ChildText("div.matchTeam.team2") + " " + element.Attr("team2")
			matchTime, _ := strconv.ParseInt(strings.TrimSpace(element.ChildAttr("div.matchTime", "data-unix")), 10, 64)
			date := time.UnixMilli(matchTime).UTC()

			m := models.UpcomingMatch{Link: link, Stars: stars, Team1: team1, Team2: team2, MatchTime: date}
			matches = append(matches, m)
		})
	})

	c.OnRequest(func(request *colly.Request) {
		request.Headers.Set("User-Agent", RandomString())
		app.log.Infof("Request to %v", request.URL.RequestURI())
	})

	c.OnError(func(r *colly.Response, err error) {
		app.log.Errorf("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c.Visit(url)
	if err != nil {
		app.log.Fatal(err)
	}

	js, err := json.MarshalIndent(matches, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	if err != nil {
		app.log.Fatal(err)
	}
}
