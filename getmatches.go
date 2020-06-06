package jpmcpvpgo

import (
	"encoding/json"
	"net/url"
	"time"
)

type Matches []struct {
	ID           string       `json:"id"`
	Map          string       `json:"map"`
	Server       string       `json:"server"`
	Gamemode     string       `json:"gamemode"`
	Ranked       bool         `json:"ranked"`
	Win          string       `json:"win"`
	KillCount    int          `json:"kill_count"`
	DeathCount   int          `json:"death_count"`
	PlayerCount  int          `json:"player_count"`
	Started      time.Time    `json:"started"`
	Finished     time.Time    `json:"finished"`
	Teams        []Teams      `json:"teams"`
	Objectives   []Objectives `json:"objectives"`
	PermalinkURL string       `json:"permalink_url"`
}
type MatchesPlayers struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	Score      int    `json:"score"`
	KillCount  int    `json:"kill_count"`
	DeathCount int    `json:"death_count"`
	ShotCount  int    `json:"shot_count"`
	HitCount   int    `json:"hit_count"`
	PaintCount int    `json:"paint_count"`
}
type Teams struct {
	Name           string           `json:"name"`
	Score          int              `json:"score"`
	KillCount      int              `json:"kill_count"`
	DeathCount     int              `json:"death_count"`
	MatchesPlayers []MatchesPlayers `json:"players"`
}
type Objectives struct {
	Player string    `json:"player"`
	Name   string    `json:"name"`
	Color  string    `json:"color"`
	Type   string    `json:"type"`
	Team   string    `json:"team"`
	Time   time.Time `json:"time"`
}

type MatchesOptions struct {
	Max_Id string
	Limit  string
	Fields string
}

func (session Session) GetMatches(options MatchesOptions) (*Matches, error) {
	values := url.Values{}
	if options.Max_Id != "" {
		values.Add("max_id", options.Max_Id)
	}
	if options.Limit != "" {
		values.Add("limit", options.Limit)
	}
	if options.Fields != "" {
		values.Add("fields", options.Fields)
	}

	body, err := session.request("matches", values)

	if err != nil {
		return nil, err
	}

	matches := Matches{}

	err = json.Unmarshal(body, &matches)

	return &matches, err
}
