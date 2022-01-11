package internal

import (
	"github.com/est5/gohltv/internal/helpers"
	"github.com/est5/gohltv/pkg/models"
	"github.com/gocolly/colly"
	"net/http"
	"strconv"
	"time"
)

func (app application) GetMatchesStats(w http.ResponseWriter, r *http.Request) {
	c := colly.NewCollector()
	url := helpers.GetUrl(r)
	var teams []models.StatsTeams
	c.Limit(&colly.LimitRule{RandomDelay: 5 * time.Second})

	c.OnHTML("tr", func(e *colly.HTMLElement) {
		link := helpers.Prefix + e.ChildAttr("a", "href")
		name := e.ChildText("a")
		statsDetail := e.ChildText("td.statsDetail")
		kddif := e.ChildText("td.kdDiffCol")
		rating, _ := strconv.ParseFloat(e.ChildText("td.ratingCol"), 64)
		println(name, kddif, statsDetail, rating)
		if rating != 0.0 {
			team := models.StatsTeams{
				Link: link,
				Name: name,
				//Maps:   statsDetail,
				KDDiff: kddif,
				//KD:     statsDetail,
				Rating: rating,
			}
			teams = append(teams, team)
			println(statsDetail)
			//os.Exit(1)
		}

	})

	helpers.Visit(w, c.Visit(url), url, teams)
}
