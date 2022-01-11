package helpers

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"io"
	"math/rand"
	"net/http"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var JsonMarshalingError = "Error Marshaling to JSON"
var UrlVisitError = "Error while visiting url"

const Prefix = "https://www.hltv.org"

func ToJson(slice interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	err := e.Encode(slice)
	if err != nil {
		return errors.New("marshaling to json error")
	}
	return nil
}

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func MatchesLink(uri string) string {
	switch uri {
	case "top":
		return "https://www.hltv.org/matches?predefinedFilter=top_tier"
	case "lan":
		return "https://www.hltv.org/matches?predefinedFilter=lan_only"
	}
	return "https://www.hltv.org/matches"
}

func Visit(w http.ResponseWriter, visit error, url string, toJson interface{}) {
	if visit != nil {
		http.Error(w, UrlVisitError, http.StatusBadRequest)
		log.Fatalf("Bad request to %v : %v ", url, visit)
		return
	}

	if err := ToJson(toJson, w); err != nil {
		http.Error(w, JsonMarshalingError, http.StatusInternalServerError)
		log.Fatal("Error marshaling to json %v", err)
		return
	}
}
