package main

import (
	"encoding/json"
	"github.com/est5/gohltv/pkg/models"
	"github.com/gocolly/colly"
	"net/http"
	"strconv"
	"strings"
)

func (app *application) GetOngoingEvents(w http.ResponseWriter, r *http.Request) {
	c := colly.NewCollector()
	prefix := "https://www.hltv.org"
	url := eventsParams(r)
	var events []models.OngoingEvent

	c.OnHTML("div#ALL.tab-content", func(e *colly.HTMLElement) {
		e.ForEach("a.a-reset.ongoing-event", func(_ int, element *colly.HTMLElement) {
			link := prefix + element.Attr("href")
			name := element.ChildText("div.event-name-small")
			date := element.ChildText("span.col-desc")
			eventLink := strings.Split(link, "/")
			eventId, _ := strconv.Atoi(eventLink[4])
			event := models.OngoingEvent{
				Link:    link,
				Name:    name,
				Date:    date,
				EventId: eventId,
			}
			events = append(events, event)
		})
	})

	err := c.Visit(url)
	if err != nil {
		app.log.Error(err)
		err := c.Visit("https://www.hltv.org/events")
		if err != nil {
			app.log.Fatal(err)
		}
	}

	c.OnRequest(func(request *colly.Request) {
		request.Headers.Set("User-Agent", RandomString())
		app.log.Infof("Request to %v", request.URL.RequestURI())
	})

	js, err := json.MarshalIndent(events, "", " ")
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
