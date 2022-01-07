package models

type UpcomingMatch struct {
	Link      string
	Stars     string
	Team1     string
	Team2     string
	MatchTime string
}

type NewsArticle struct {
	Link          string
	Text          string
	CommentsCount int
	Date          string
}

type LiveMatch struct {
}

type ResultSet struct {
	Link        string
	ResultScore string
	Team1       string
	Team2       string
	MatchTime   string
	Map         string
}
