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
	url := matchesLink(vars["type"])
	var matches []models.UpcomingMatch

	c.OnHTML("div.upcomingMatchesContainer", func(e *colly.HTMLElement) {
		e.ForEach("div.upcomingMatch", func(i int, element *colly.HTMLElement) {
			link := "https://www.hltv.org" + element.ChildAttr("a.match", "href")
			stars := element.Attr("stars")
			team1 := element.ChildText("div.matchTeam.team1")
			team1id, _ := strconv.Atoi(element.Attr("team1"))
			team2 := element.ChildText("div.matchTeam.team2")
			team2id, _ := strconv.Atoi(element.Attr("team2"))
			matchTime, _ := strconv.ParseInt(
				strings.TrimSpace(element.ChildAttr("div.matchTime", "data-unix")),
				10,
				64,
			)
			date := time.UnixMilli(matchTime).UTC()

			m := models.UpcomingMatch{
				Link:      link,
				Stars:     stars,
				Team1:     team1,
				Team1Id:   team1id,
				Team2:     team2,
				Team2Id:   team2id,
				MatchTime: date.Format("2 Jan 06, 15:04UTC"),
			}
			matches = append(matches, m)
		})
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
	_, err = w.Write(js)
	if err != nil {
		app.log.Fatal(err)
	}
}
