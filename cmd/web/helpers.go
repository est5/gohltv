package main

import (
	"math/rand"
	"net/http"
	"net/url"
	"strings"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func matchesLink(uri string) string {
	switch uri {
	case "top":
		return "https://www.hltv.org/matches?predefinedFilter=top_tier"
	case "lan":
		return "https://www.hltv.org/matches?predefinedFilter=lan_only"
	}
	return "https://www.hltv.org/matches"
}

func resultsParams(r *http.Request) (url string) {
	params := r.URL.Query()
	if len(params) == 0 {
		return "https://www.hltv.org/results/"
	}
	url = "https://www.hltv.org/results/"
	url = getParam(&params, &url, "stars")
	url = getDate(&params, &url)
	url = getParam(&params, &url, "offset")
	url = getParam(&params, &url, "matchType")
	url = getParam(&params, &url, "content")
	url = getParam(&params, &url, "map")
	url = getParam(&params, &url, "gameType")
	return url
}

func getDate(params *url.Values, url *string) string {
	if (params.Get("startDate") != "") && (params.Get("endDate") != "") {
		if strings.LastIndex(*url, "?") != -1 {
			*url += "&startDate=" + params.Get("startDate") + "&endDate=" + params.Get("endDate")
		} else {
			*url += "?startDate=" + params.Get("startDate") + "&endDate=" + params.Get("endDate")
		}
	}
	return *url
}

func getParam(params *url.Values, url *string, paramName string) string {
	if params.Get(paramName) != "" {
		if strings.LastIndex(*url, "?") != -1 {
			*url += "&" + paramName + "=" + params.Get(paramName)
		} else {
			*url += "?" + paramName + "=" + params.Get(paramName)
		}
	}
	return *url
}
