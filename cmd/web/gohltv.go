package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

type application struct {
	log *log.Logger
}

func main() {
	app := application{log: log.New()}
	srv := &http.Server{Addr: ":3000", Handler: app.routes()}
	log.Infoln("Starting at port :3000")
	srv.ListenAndServe()

}
