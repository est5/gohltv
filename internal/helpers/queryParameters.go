package helpers

import (
	"net/http"
	"net/url"
	"strings"
)

const (
	eventType = "eventType"
	prizeMin  = "prizeMin"
	prizeMax  = "prizeMax"
	team      = "team"
	player    = "player"
	startDate = "startDate"
	endDate   = "endDate"
	gameType  = "gameType"
	offset    = "offset"
	stars     = "stars"
	content   = "content"
	matchType = "matchType"
	gameMap   = "map"
	event     = "event"
)

func ResultsParams(r *http.Request) (url string) {
	params := r.URL.Query()
	const resultsDefaultUrl = "https://www.hltv.org/results/"
	if len(params) == 0 {
		return resultsDefaultUrl
	}
	url = resultsDefaultUrl
	url = getParam(&params, &url, stars)
	url = getRangeParam(&params, &url, startDate, endDate)
	url = getParam(&params, &url, offset)
	url = getParam(&params, &url, matchType)
	url = getParam(&params, &url, content)
	url = getParam(&params, &url, gameMap)
	url = getParam(&params, &url, gameType)
	url = getParam(&params, &url, team)
	url = getParam(&params, &url, player)
	url = getParam(&params, &url, event)
	return url
}

func EventsParams(r *http.Request) (url string) {
	params := r.URL.Query()
	const eventsDefaultUrl = "https://www.hltv.org/events#tab-ALL"
	if len(params) == 0 {
		return eventsDefaultUrl
	}
	url = "https://www.hltv.org/events/"
	url = getParam(&params, &url, eventType)
	url = getRangeParam(&params, &url, prizeMin, prizeMax)
	url = getParam(&params, &url, team)   // id
	url = getParam(&params, &url, player) // id

	const tabPostfix = "#tab-ALL"
	return url + tabPostfix
}

func EventsArchiveParams(r *http.Request) (url string) {
	params := r.URL.Query()
	const archiveUrl = "https://www.hltv.org/events/archive/"
	if len(params) == 0 {
		return archiveUrl
	}
	url = archiveUrl

	url = getParam(&params, &url, eventType)
	url = getRangeParam(&params, &url, prizeMin, prizeMax)
	url = getParam(&params, &url, team)   // id
	url = getParam(&params, &url, player) // id
	url = getRangeParam(&params, &url, startDate, endDate)
	url = getParam(&params, &url, gameType)
	url = getParam(&params, &url, offset)

	return url
}

func getRangeParam(params *url.Values, url *string, start, end string) string {
	if (params.Get(start) != "") && (params.Get(end) != "") {
		if strings.LastIndex(*url, "?") != -1 {
			*url += "&" + start + "=" + params.Get(start) + "&" + end + "=" + params.Get(end)
		} else {
			*url += "?" + start + "=" + params.Get(start) + "&" + end + "=" + params.Get(end)
		}
	}
	return *url
}

func getParam(params *url.Values, url *string, paramName string) string {
	for _, val := range *params {
		for i := 0; i < len(val); i++ {
			if params.Get(paramName) != "" {
				if strings.LastIndex(*url, "?") != -1 {
					*url += "&" + paramName + "=" + val[i]
				} else {
					*url += "?" + paramName + "=" + val[i]
				}
			}
		}
	}
	return *url
}
