package main

import (
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
		app.log.Errorf("Bad request to %v", url)
		http.Error(w, "marshaling to json error", http.StatusBadRequest)
		return
	}
	if err := ToJson(events, w); err != nil {
		app.log.Errorf("Error marshaling to json %v", err)
		http.Error(w, "marshaling to json error", http.StatusInternalServerError)
		return
	}

}

func (app application) GetUpcomingEvents(w http.ResponseWriter, r *http.Request) {
	c := colly.NewCollector()
	prefix := "https://www.hltv.org"
	url := eventsParams(r)
	var events []models.UpcomingEvent

	c.OnHTML("div.events-month", func(e *colly.HTMLElement) {
		e.ForEach("a.a-reset.small-event.standard-box", func(_ int, element *colly.HTMLElement) {
			link := prefix + element.Attr("href")
			eventId := strings.Split(link, "/")
			id, _ := strconv.Atoi(eventId[4])
			name := element.ChildText("div.text-ellipsis")
			allLines := element.ChildText("td.col-value.small-col")
			teamNum := allLines[:strings.IndexAny(allLines, "$O")]
			prize := element.ChildAttr("td.col-value.small-col.prizePoolEllipsis", "title")
			country := element.ChildText("span.smallCountry")
			date := element.ChildText("span.col-desc")
			date = strings.Split(date, "|")[1]
			country = strings.TrimSpace(strings.Trim(country, "|"))

			event := models.UpcomingEvent{
				Link:          strings.TrimSpace(link),
				Name:          strings.TrimSpace(name),
				EventId:       id,
				Date:          strings.TrimSpace(date),
				Prize:         prize,
				NumberOfTeams: teamNum,
				EventLocation: country,
			}

			events = append(events, event)
		})

		e.ForEach("a.a-reset.standard-box.big-event", func(_ int, element *colly.HTMLElement) {
			link := prefix + element.Attr("href")
			eventId := strings.Split(link, "/")
			id, _ := strconv.Atoi(eventId[4])
			name := element.ChildText("div.big-event-name")
			location := element.ChildText("div.location-top-teams")
			date := element.ChildText("td.col-value.col-date")
			additionalInfo := strings.Split(element.ChildText("div.additional-info"), "\n")
			prize := strings.TrimSpace(additionalInfo[1])
			teams := strings.TrimSpace(additionalInfo[2])

			event := models.UpcomingEvent{
				Link:          link,
				Name:          name,
				EventId:       id,
				Date:          date,
				Prize:         prize,
				NumberOfTeams: teams,
				EventLocation: location,
			}

			events = append(events, event)
		})
	})

	err := c.Visit(url)
	if err != nil {
		app.log.Errorf("Bad request to %v", url)
		http.Error(w, "marshaling to json error", http.StatusBadRequest)
		return
	}
	if err := ToJson(events, w); err != nil {
		app.log.Errorf("Error marshaling to json %v", err)
		http.Error(w, "marshaling to json error", http.StatusInternalServerError)
		return
	}

}
