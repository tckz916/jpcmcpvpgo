package jpmcpvpgo

import (
	"encoding/json"
	"net/url"
)

type PlayingMaps []struct {
	ID           string              `json:"id"`
	Name         string              `json:"name"`
	Version      string              `json:"version"`
	Proto        string              `json:"proto"`
	Authors      []PlayingMapAuthors `json:"authors"`
	Objective    string              `json:"objective"`
	Teams        []Teams             `json:"teams"`
	Gamemodes    []string            `json:"gamemodes"`
	Type         string              `json:"type"`
	Rotations    []string            `json:"rotations"`
	Path         string              `json:"path"`
	HasImage     bool                `json:"has_image"`
	Rates        []PlayingMapRates   `json:"rates"`
	Servers      []PlayingMapServers `json:"servers"`
	PermalinkURL string              `json:"permalink_url"`
}
type PlayingMapAuthors struct {
	Name         string `json:"name"`
	UUID         string `json:"uuid"`
	Contribution string `json:"contribution"`
}
type PlayingMapTeams struct {
	Name  string `json:"name"`
	Max   int    `json:"max"`
	Color string `json:"color"`
}
type PlayingMapRate struct {
	Num1 int `json:"1"`
	Num2 int `json:"2"`
	Num3 int `json:"3"`
	Num4 int `json:"4"`
	Num5 int `json:"5"`
}
type PlayingMapRates struct {
	Version string         `json:"version"`
	Total   int            `json:"total"`
	Rate    float64        `json:"rate"`
	Rates   PlayingMapRate `json:"rates"`
}
type PlayingMapServers struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Region     string `json:"region"`
	CurrentMap string `json:"current_map"`
	NextMap    string `json:"next_map"`
}

func (session Session) GetPlayingMaps() (*PlayingMaps, error) {
	values := url.Values{}
	values.Add("playing", "1")

	body, err := session.request("maps", values)

	if err != nil {
		return nil, err
	}

	playingMaps := PlayingMaps{}

	err = json.Unmarshal(body, &playingMaps)

	return &playingMaps, err
}