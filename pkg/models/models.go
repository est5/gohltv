package models

type UpcomingMatch struct {
	Link      string `json:"link,omitempty"`
	Stars     string `json:"stars,omitempty"`
	Team1     string `json:"team-1,omitempty"`
	Team1Id   int    `json:"team-1-id,omitempty"`
	Team2     string `json:"team-2,omitempty"`
	Team2Id   int    `json:"team-2-id,omitempty"`
	MatchTime string `json:"match-time,omitempty"`
}

type NewsArticle struct {
	Link          string `json:"link,omitempty"`
	Text          string `json:"text,omitempty"`
	CommentsCount int    `json:"comments-count,omitempty"`
	Date          string `json:"date,omitempty"`
}

type LiveMatch struct {
}

type ResultSet struct {
	Link        string `json:"link,omitempty"`
	ResultScore string `json:"result-score,omitempty"`
	Team1       string `json:"team-1,omitempty"`
	Team2       string `json:"team-2,omitempty"`
	MatchTime   string `json:"match-time,omitempty"`
	Map         string `json:"map,omitempty"`
}

type OngoingEvent struct {
	Link    string `json:"link,omitempty"`
	Name    string `json:"name,omitempty"`
	EventId int    `json:"event-id,omitempty"`
	Date    string `json:"date,omitempty"`
}

type UpcomingEvent struct {
	Link    string `json:"link,omitempty"`
	Name    string `json:"name,omitempty"`
	EventId int    `json:"event-id,omitempty"`
	Date    string `json:"date,omitempty"`
	Prize   string `json:"prize,omitempty"`
	Type    string `json:"type,omitempty"`
	Teams   string `json:"teams,omitempty"`
	Country string `json:"country,omitempty"`
}
