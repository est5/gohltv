package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type application struct {
	log *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	app := application{log: log.New()}
	srv := &http.Server{Addr: *addr, Handler: app.routes()}
	log.Infof("Starting serve on %s\n", *addr)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
