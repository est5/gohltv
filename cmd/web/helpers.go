package main

import (
	"math/rand"
	"net/http"
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
	if params.Get("stars") != "" {
		if strings.LastIndex(url, "?") != -1 {
			url += "&stars=" + params.Get("stars")
		} else {
			url += "?stars=" + params.Get("stars")
		}
	}
	if params.Get("offset") != "" {
		if strings.LastIndex(url, "?") != -1 {
			url += "&offset=" + params.Get("offset")
		} else {
			url += "?offset=" + params.Get("offset")
		}
	}
	if params.Get("startDate") != "" {
		if strings.LastIndex(url, "?") != -1 {
			url += "&startDate=" + params.Get("startDate")
		} else {
			url += "?startDate=" + params.Get("startDate")
		}
	}

	return url
}
