package jpmcpvpgo

import (
	"encoding/json"
	"net/url"
	"time"
)

type Matches []Match

type Match struct {
	ID              *string           `json:"id,omitempty"`
	Map             *string           `json:"map,omitempty"`
	Server          *string           `json:"server,omitempty"`
	GameMode        *string           `json:"gamemode,omitempty"`
	Ranked          *bool             `json:"ranked,omitempty"`
	Win             *string           `json:"win,omitempty"`
	KillCount       *int              `json:"kill_count,omitempty"`
	DeathCount      *int              `json:"death_count,omitempty"`
	PlayerCount     *int              `json:"player_count,omitempty"`
	Started         *time.Time        `json:"started,omitempty"`
	Finished        *time.Time        `json:"finished,omitempty"`
	MatchTeams      *[]MatchTeam      `json:"teams,omitempty"`
	PermalinkURL    *string           `json:"permalink_url,omitempty"`
	MatchObjectives *[]MatchObjective `json:"objectives,omitempty,omitempty"`
}

type MatchTeam struct {
	Name         *string        `json:"name,omitempty"`
	Score        *int           `json:"score,omitempty"`
	KillCount    *int           `json:"kill_count,omitempty"`
	DeathCount   *int           `json:"death_count,omitempty"`
	MatchPlayers *[]MatchPlayer `json:"players,omitempty"`
}

type MatchPlayer struct {
	UUID       *string `json:"uuid,omitempty"`
	Name       *string `json:"name,omitempty"`
	Score      *int    `json:"score,omitempty"`
	KillCount  *int    `json:"kill_count,omitempty"`
	DeathCount *int    `json:"death_count,omitempty"`
	ShotCount  *int    `json:"shot_count,omitempty"`
	HitCount   *int    `json:"hit_count,omitempty"`
	PaintCount *int    `json:"paint_count,omitempty"`
}

type MatchObjective struct {
	Player *string    `json:"player,omitempty"`
	Name   *string    `json:"name,omitempty"`
	Color  *string    `json:"color,omitempty"`
	Type   *string    `json:"type,omitempty"`
	Team   *string    `json:"team,omitempty"`
	Time   *time.Time `json:"time,omitempty"`
}

type MatchesOptions struct {
	Max_Id *string
	Limit  *string
	Fields *string
}

func (session Session) GetMatches(options MatchesOptions) (*Matches, error) {
	values := url.Values{}
	if options.Max_Id != nil {
		values.Add("max_id", *options.Max_Id)
	}
	if options.Limit != nil {
		values.Add("limit", *options.Limit)
	}
	if options.Fields != nil {
		values.Add("fields", *options.Fields)
	}

	body, err := session.request("matches", values)

	if err != nil {
		return nil, err
	}

	matches := Matches{}

	err = json.Unmarshal(body, &matches)

	return &matches, err
}
