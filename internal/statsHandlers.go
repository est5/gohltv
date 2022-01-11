package internal

import (
	"github.com/est5/gohltv/internal/helpers"
	"github.com/est5/gohltv/pkg/models"
	"github.com/gocolly/colly"
	"net/http"
	"strconv"
	"strings"
)

func (app *application) GetStatsPlayers(w http.ResponseWriter, r *http.Request) {
	c := *colly.NewCollector()
	var players []models.StatsPlayer
	url := "https://www.hltv.org/stats/players" //call to helpers to get query params

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

	helpers.Visit(w, c.Visit(url), url, players)
}

func (app application) GetStatsPlayersFlashes(w http.ResponseWriter, r *http.Request) {
	c := colly.NewCollector()
	url := "https://www.hltv.org/stats/players/flashbangs" // call to helper to get params
	var playersFlashes []models.StatsPlayerFlashes

	c.OnHTML("tbody", func(e *colly.HTMLElement) {

		e.ForEach("tr", func(_ int, element *colly.HTMLElement) {
			var blob string
			link := helpers.Prefix + e.ChildAttr("a", "href")
			blob += element.Text
			splited := strings.Split(strings.TrimSpace(blob), "\n")
			name := strings.TrimSpace(splited[0])
			mapsCount, _ := strconv.Atoi(strings.TrimSpace(splited[1]))
			rounds, _ := strconv.Atoi(strings.TrimSpace(splited[2]))
			thrown, _ := strconv.ParseFloat(strings.TrimSpace(splited[3]), 64)
			FA, _ := strconv.ParseFloat(strings.TrimSpace(splited[7]), 64)
			success, _ := strconv.ParseFloat(strings.TrimSpace(splited[8]), 64)

			player := models.StatsPlayerFlashes{
				Link:        link,
				Name:        name,
				MapsCount:   mapsCount,
				RoundsCount: rounds,
				Thrown:      thrown,
				Blinded:     strings.TrimSpace(splited[4]),
				OppFlashed:  strings.TrimSpace(splited[5]),
				Diff:        strings.TrimSpace(splited[6]),
				FA:          FA,
				Success:     success,
			}

			playersFlashes = append(playersFlashes, player)
		})

	})

	helpers.Visit(w, c.Visit(url), url, playersFlashes)
}
