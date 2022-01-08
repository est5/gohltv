package models

type UpcomingMatch struct {
	Link      string
	Stars     string
	Team1     string
	Team1Id   int
	Team2     string
	Team2Id   int
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

type OngoingEvent struct {
	Link    string
	Name    string
	EventId int
	Date    string
}

type UpcomingEvent struct {
	Link    string
	Name    string
	EventId int
	Date    string
	Prize   string
	Type    string
	Teams   string
	Country string
}
