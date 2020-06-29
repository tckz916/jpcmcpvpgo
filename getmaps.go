package jpmcpvpgo

import (
	"encoding/json"
	"net/url"
)

type Maps []Map

type Map struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Version      string       `json:"version"`
	Proto        string       `json:"proto"`
	Authors      []MapAuthors `json:"authors"`
	Objective    string       `json:"objective"`
	Teams        []MapTeams   `json:"teams"`
	Gamemodes    []string     `json:"gamemodes"`
	Type         string       `json:"type"`
	Rotations    []string     `json:"rotations"`
	Path         string       `json:"path"`
	HasImage     bool         `json:"has_image"`
	Rates        []MapRates   `json:"rates"`
	PermalinkURL string       `json:"permalink_url"`
}
type MapAuthors struct {
	Name         string `json:"name"`
	UUID         string `json:"uuid"`
	Contribution string `json:"contribution"`
}
type MapTeams struct {
	Name  string `json:"name"`
	Max   int    `json:"max"`
	Color string `json:"color"`
}
type MapRate struct {
	Num1 int `json:"1"`
	Num2 int `json:"2"`
	Num3 int `json:"3"`
	Num4 int `json:"4"`
	Num5 int `json:"5"`
}
type MapRates struct {
	Version string  `json:"version"`
	Total   int     `json:"total"`
	Rate    float64 `json:"rate"`
	Rates   MapRate `json:"rates"`
}

type MapsOptions struct {
	Max_Id string
	Limit  string
}

func (session Session) GetMaps(options MapsOptions) (*Maps, error) {
	values := url.Values{}
	if options.Max_Id != "" {
		values.Add("max_id", options.Max_Id)
	}
	if options.Limit != "" {
		values.Add("limit", options.Limit)
	}

	body, err := session.request("maps", values)

	if err != nil {
		return nil, err
	}

	maps := Maps{}

	err = json.Unmarshal(body, &maps)

	return &maps, err
}
