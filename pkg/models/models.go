package models

import "time"

type UpcomingMatch struct {
	Link      string
	Stars     string
	Team1     string
	Team2     string
	MatchTime time.Time
}

type LiveMatch struct {
}