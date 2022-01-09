package internal

import (
	"github.com/est5/gohltv/internal/helpers"
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
	url := helpers.MatchesLink(vars["type"])
	var matches []models.UpcomingMatch

	c.OnHTML("div.upcomingMatchesContainer", func(e *colly.HTMLElement) {
		e.ForEach("div.upcomingMatch", func(i int, element *colly.HTMLElement) {
			link := helpers.Prefix + element.ChildAttr("a.match", "href")
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
		app.log.Errorf("Bad request to %v", url)
		http.Error(w, helpers.UrlVisitError, http.StatusBadRequest)
		return
	}

	if err := helpers.ToJson(matches, w); err != nil {
		app.log.Errorf("Error marshaling to json %v", err)
		http.Error(w, helpers.JsonMarshalingError, http.StatusInternalServerError)
		return
	}

}

func (app *application) GetLiveMatches(w http.ResponseWriter, r *http.Request) {
	c := colly.NewCollector()
	//vars := mux.Vars(r)
	url := helpers.Prefix + "/matches"
	var liveMatches []models.LiveMatch
	c.OnHTML("div.liveMatch-container", func(e *colly.HTMLElement) {
		link := helpers.Prefix + e.ChildAttr("a.match.a-reset", "href")
		stars, _ := strconv.Atoi(e.Attr("stars"))
		matchId, _ := strconv.Atoi(e.Attr("data-scorebot-id"))
		team1Id, _ := strconv.Atoi(e.Attr("team1"))
		team2Id, _ := strconv.Atoi(e.Attr("team2"))
		maps := e.Attr("data-maps")
		teams := e.ChildAttrs("img.matchTeamLogo", "alt")
		team1 := teams[0]
		team2 := teams[1]
		matchEvent := e.ChildAttr("img.matchEventLogo", "alt")
		matchType := e.ChildText("div.matchMeta")
		liveMatch := models.LiveMatch{
			Link:           link,
			MatchStars:     stars,
			MatchId:        matchId,
			Maps:           maps,
			Team1:          team1,
			Team1Id:        team1Id,
			Team2:          team2,
			Team2Id:        team2Id,
			MatchEventName: matchType,
			MatchType:      matchEvent,
		}

		liveMatches = append(liveMatches, liveMatch)
	})

	err := c.Visit(url)
	if err != nil {
		app.log.Errorf("Bad request to %v", url)
		http.Error(w, helpers.UrlVisitError, http.StatusBadRequest)
		return
	}

	if err := helpers.ToJson(liveMatches, w); err != nil {
		app.log.Errorf("Error marshaling to json %v", err)
		http.Error(w, helpers.JsonMarshalingError, http.StatusInternalServerError)
		return
	}
}
