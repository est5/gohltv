package main

import (
	"encoding/json"
	"github.com/est5/gohltv/pkg/models"
	"github.com/gocolly/colly"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (app *application) GetResults(w http.ResponseWriter, r *http.Request) {
	c := colly.NewCollector()
	var results []models.ResultSet
	prefix := "https://www.hltv.org"
	url := resultsParams(r)
	c.OnHTML("div.results", func(e *colly.HTMLElement) {
		e.ForEach("div.result-con", func(_ int, element *colly.HTMLElement) {
			link := prefix + element.ChildAttr("a.a-reset", "href")
			matchTime, _ := strconv.ParseInt(
				strings.TrimSpace(element.Attr("data-zonedgrouping-entry-unix")),
				10,
				64,
			)
			date := time.UnixMilli(matchTime).UTC()
			layout := "2 Jan 06, 15:04UTC"
			formattedDate := date.Format(layout)
			team1 := element.ChildText("div.line-align.team1")
			team2 := element.ChildText("div.line-align.team2")
			mapText := element.ChildText("div.map-text")
			resultScore := element.ChildText("td.result-score")
			result := models.ResultSet{
				Link:        link,
				ResultScore: resultScore,
				Team1:       team1,
				Team2:       team2,
				MatchTime:   formattedDate,
				Map:         mapText,
			}

			results = append(results, result)
		})
	})

	err := c.Visit(url)
	if err != nil {
		app.log.Error(err)
		err := c.Visit("https://www.hltv.org/results")
		if err != nil {
			app.log.Fatal(err)
		}
	}

	js, err := json.MarshalIndent(results, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(js)
	if err != nil {
		app.log.Fatal(err)
	}
}
