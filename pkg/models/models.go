package models

type UpcomingMatch struct {
	Link      string `json:"link,omitempty"`
	Stars     string `json:"stars,omitempty"`
	Team1     string `json:"team1,omitempty"`
	Team1Id   int    `json:"team1Id,omitempty"`
	Team2     string `json:"team2,omitempty"`
	Team2Id   int    `json:"team2Id,omitempty"`
	MatchTime string `json:"matchTime,omitempty"`
}

type NewsArticle struct {
	Link          string `json:"link,omitempty"`
	Text          string `json:"text,omitempty"`
	CommentsCount int    `json:"commentsCount,omitempty"`
	Date          string `json:"date,omitempty"`
}

type LiveMatch struct {
}

type ResultSet struct {
	Link        string `json:"link,omitempty"`
	ResultScore string `json:"resultScore,omitempty"`
	Team1       string `json:"team1,omitempty"`
	Team2       string `json:"team2,omitempty"`
	MatchTime   string `json:"matchTime,omitempty"`
	Map         string `json:"map,omitempty"`
}

type OngoingEvent struct {
	Link    string `json:"link,omitempty"`
	Name    string `json:"name,omitempty"`
	EventId int    `json:"eventId,omitempty"`
	Date    string `json:"date,omitempty"`
}

type UpcomingEvent struct {
	Link          string `json:"link,omitempty"`
	Name          string `json:"name,omitempty"`
	EventId       int    `json:"eventId,omitempty"`
	Date          string `json:"date,omitempty"`
	Prize         string `json:"prize,omitempty"`
	NumberOfTeams string `json:"numberOfTeams,omitempty"`
	EventLocation string `json:"eventLocation,omitempty"`
}
