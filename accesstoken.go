package jpmcpvpgo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var tokenURL string = "https://minecraft.jp/oauth/token"

type TokenRequest struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type ClientCredentials struct {
	AccessToken string      `json:"access_token"`
	ExpiresIn   int         `json:"expires_in"`
	TokenType   string      `json:"token_type"`
	Scope       interface{} `json:"scope"`
}

func GetAccessToken(client_id string, client_secret string) (client *ClientCredentials, err error) {
	tokenRequest := TokenRequest{"client_credentials", client_id, client_secret}
	b, err := json.Marshal(tokenRequest)

	if err != nil {
		return nil, err
	}

	resp, err := http.Post(tokenURL, "application/json", bytes.NewBuffer(b))

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	client = &ClientCredentials{}

	err = json.Unmarshal(body, &client)

	return client, err
}
