package helpers

import "net/http"

func GetUrl(r *http.Request) string {
	return Prefix + r.RequestURI
}
