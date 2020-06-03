package jpmcpvpgo

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type PlayersOptions struct {
	Id     string
	Fields string
}

type Players struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	UUID            string    `json:"uuid"`
	LastLogin       time.Time `json:"last_login"`
	LastLoginServer string    `json:"last_login_server"`
	LastLogout      time.Time `json:"last_logout"`
	Banned          bool      `json:"banned"`
	Matches         MatchesData `json:"matches"`
	PermalinkURL    string    `json:"permalink_url"`
}

type MatchesData []struct {
	DeathCount              int       `json:"death_count"`
	EnvdeathCount           int       `json:"envdeath_count"`
	Finished                time.Time `json:"finished"`
	Gamemode                string    `json:"gamemode"`
	GlickoKillDeviationDiff float64   `json:"glicko_kill_deviation_diff,omitempty"`
	GlickoKillRatingDiff    float64   `json:"glicko_kill_rating_diff,omitempty"`
	ID                      string    `json:"id"`
	KillCount               int       `json:"kill_count"`
	Map                     string    `json:"map"`
	Result                  string    `json:"result"`
	Score                   int       `json:"score"`
	Started                 time.Time `json:"started"`
	Time                    time.Time `json:"time"`
}

func (session Session) GetPlayers(options PlayersOptions) (*Players, error) {
	if options.Id == "" {
		return nil, fmt.Errorf("idがないよ！")
	}
	values := url.Values{}
	if options.Fields != "" {
		values.Add("fields", options.Fields)
	}

	body, err := session.request("players/"+options.Id, values)

	if err != nil {
		return nil, err
	}

	players := Players{}

	err = json.Unmarshal(body, &players)

	return &players, err
}
