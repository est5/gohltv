package helpers

import (
	"net/http"
	"net/url"
	"strings"
)

func ResultsParams(r *http.Request) (url string) {
	params := r.URL.Query()
	if len(params) == 0 {
		return "https://www.hltv.org/results/"
	}
	url = "https://www.hltv.org/results/"
	url = getParam(&params, &url, "stars")
	url = getRangeParam(&params, &url, "startDate", "endDate")
	url = getParam(&params, &url, "offset")
	url = getParam(&params, &url, "matchType")
	url = getParam(&params, &url, "content")
	url = getParam(&params, &url, "map")
	url = getParam(&params, &url, "gameType")
	url = getParam(&params, &url, "team")
	url = getParam(&params, &url, "player")
	url = getParam(&params, &url, "event")
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

func EventsParams(r *http.Request) (url string) {
	params := r.URL.Query()
	if len(params) == 0 {
		return "https://www.hltv.org/events#tab-ALL"
	}
	url = "https://www.hltv.org/events/"
	url = getParam(&params, &url, "eventType")
	url = getRangeParam(&params, &url, "prizeMin", "prizeMax")
	url = getParam(&params, &url, "team")   // id
	url = getParam(&params, &url, "player") // id

	return url + "#tab-ALL"
}

func EventsArchiveParams(r *http.Request) (url string) {
	params := r.URL.Query()
	if len(params) == 0 {
		return "https://www.hltv.org/events/archive/"
	}
	url = "https://www.hltv.org/events/archive/"

	url = getParam(&params, &url, "eventType")
	url = getRangeParam(&params, &url, "prizeMin", "prizeMax")
	url = getParam(&params, &url, "team")   // id
	url = getParam(&params, &url, "player") // id
	url = getRangeParam(&params, &url, "startDate", "endDate")
	url = getParam(&params, &url, "gameType")
	url = getParam(&params, &url, "offset")

	return url
}
