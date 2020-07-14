package jpmcpvpgo

import (
	"encoding/json"
	"net/url"
)

type PlayingMaps []PlayingMap

type PlayingMap struct {
	ID           *string              `json:"id,omitempty"`
	Name         *string              `json:"name,omitempty"`
	Version      *string              `json:"version,omitempty"`
	Proto        *string              `json:"proto,omitempty"`
	Authors      *[]PlayingMapAuthors `json:"authors,omitempty"`
	Objective    *string              `json:"objective,omitempty"`
	Teams        *[]PlayingMapTeams   `json:"teams,omitempty"`
	Gamemodes    *[]string            `json:"gamemodes,omitempty"`
	Type         *string              `json:"type,omitempty"`
	Rotations    *[]string            `json:"rotations,omitempty"`
	Path         *string              `json:"path,omitempty"`
	HasImage     *bool                `json:"has_image,omitempty"`
	Rates        *[]PlayingMapRates   `json:"rates,omitempty"`
	Servers      *[]PlayingMapServers `json:"servers,omitempty"`
	PermalinkURL *string              `json:"permalink_url,omitempty"`
}
type PlayingMapAuthors struct {
	Name         *string `json:"name,omitempty"`
	UUID         *string `json:"uuid,omitempty"`
	Contribution *string `json:"contribution,omitempty"`
}
type PlayingMapTeams struct {
	Name  *string `json:"name,omitempty"`
	Max   *int    `json:"max,omitempty"`
	Color *string `json:"color,omitempty"`
}
type PlayingMapRate struct {
	Num1 *int `json:"1,omitempty"`
	Num2 *int `json:"2,omitempty"`
	Num3 *int `json:"3,omitempty"`
	Num4 *int `json:"4,omitempty"`
	Num5 *int `json:"5,omitempty"`
}
type PlayingMapRates struct {
	Version *string         `json:"version,omitempty"`
	Total   *int            `json:"total,omitempty"`
	Rate    *float64        `json:"rate,omitempty"`
	Rates   *PlayingMapRate `json:"rates,omitempty"`
}
type PlayingMapServers struct {
	ID         *string `json:"id,omitempty"`
	Name       *string `json:"name,omitempty"`
	Region     *string `json:"region,omitempty"`
	CurrentMap *string `json:"current_map,omitempty"`
	NextMap    *string `json:"next_map,omitempty"`
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
