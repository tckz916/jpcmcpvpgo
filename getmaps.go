package jpmcpvpgo

import (
	"encoding/json"
	"net/url"
)

type Maps []Map

type Map struct {
	ID           *string       `json:"id,omitempty"`
	Name         *string       `json:"name,omitempty"`
	Version      *string       `json:"version,omitempty"`
	Proto        *string       `json:"proto,omitempty"`
	Authors      *[]MapAuthors `json:"authors,omitempty"`
	Objective    *string       `json:"objective,omitempty"`
	Teams        *[]MapTeams   `json:"teams,omitempty"`
	Gamemodes    *[]string     `json:"gamemodes,omitempty"`
	Type         *string       `json:"type,omitempty"`
	Rotations    *[]string     `json:"rotations,omitempty"`
	Path         *string       `json:"path,omitempty"`
	HasImage     *bool         `json:"has_image,omitempty"`
	Rates        *[]MapRates   `json:"rates,omitempty"`
	PermalinkURL *string       `json:"permalink_url,omitempty"`
}
type MapAuthors struct {
	Name         *string `json:"name,omitempty"`
	UUID         *string `json:"uuid,omitempty"`
	Contribution *string `json:"contribution,omitempty"`
}
type MapTeams struct {
	Name  *string `json:"name,omitempty"`
	Max   *int    `json:"max,omitempty"`
	Color *string `json:"color,omitempty"`
}
type MapRate struct {
	Num1 *int `json:"1,omitempty"`
	Num2 *int `json:"2,omitempty"`
	Num3 *int `json:"3,omitempty"`
	Num4 *int `json:"4,omitempty"`
	Num5 *int `json:"5,omitempty"`
}
type MapRates struct {
	Version *string  `json:"version,omitempty"`
	Total   *int     `json:"total,omitempty"`
	Rate    *float64 `json:"rate,omitempty"`
	Rates   *MapRate `json:"rates,omitempty"`
}

type MapsOptions struct {
	Max_Id *string
	Limit  *string
}

func (session Session) GetMaps(options MapsOptions) (*Maps, error) {
	values := url.Values{}
	if options.Max_Id != nil {
		values.Add("max_id", *options.Max_Id)
	}
	if options.Limit != nil {
		values.Add("limit", *options.Limit)
	}

	body, err := session.request("maps", values)

	if err != nil {
		return nil, err
	}

	maps := Maps{}

	err = json.Unmarshal(body, &maps)

	return &maps, err
}
