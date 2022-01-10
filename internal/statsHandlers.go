package internal

import (
	"github.com/est5/gohltv/internal/helpers"
	"github.com/est5/gohltv/pkg/models"
	"github.com/gocolly/colly"
	"net/http"
	"strconv"
	"strings"
)

// stats page
func (app application) GetStatsPlayers(w http.ResponseWriter, r *http.Request) {
	c := colly.NewCollector()
	var players []models.StatsPlayer
	url := "https://www.hltv.org/stats/players"

	c.OnHTML("tr", func(e *colly.HTMLElement) {
		link := helpers.Prefix + e.ChildAttr("a", "href")
		name := strings.Split(e.ChildAttr("a", "href"), "/")
		var teams []string
		teams = append(teams, e.ChildAttrs("img.logo", "title")...)
		var n string
		var statsDetail []string
		e.ForEach("td.statsDetail", func(_ int, element *colly.HTMLElement) {
			statsDetail = append(statsDetail, element.Text)
		})
		rounds, _ := strconv.Atoi(e.ChildText("td.statsDetail.gtSmartphone-only"))
		rating, _ := strconv.ParseFloat(e.ChildText("td.ratingCol"), 64)
		var maps, kdDiff int
		var kd float64

		if len(name) > 1 {
			n = name[4]
			maps, _ = strconv.Atoi(statsDetail[0])
			kdDiff, _ = strconv.Atoi(statsDetail[1])
			kd, _ = strconv.ParseFloat(statsDetail[2], 64)
		}

		player := models.StatsPlayer{
			Link:        link,
			Name:        n,
			Team:        teams,
			MapsCount:   maps,
			RoundsCount: rounds,
			KDDiff:      kdDiff,
			KD:          kd,
			Rating:      rating,
		}
		players = append(players, player)

	})

	err := c.Visit(url)
	if err != nil {
		app.log.Errorf("Bad request to %v", url)
		http.Error(w, helpers.UrlVisitError, http.StatusBadRequest)
		return
	}

	if err := helpers.ToJson(players, w); err != nil {
		app.log.Errorf("Error marshaling to json %v", err)
		http.Error(w, helpers.JsonMarshalingError, http.StatusInternalServerError)
		return
	}
}

// player stats page

// team stats page
