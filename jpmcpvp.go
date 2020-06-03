package jpmcpvpgo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var baseURL string = "https://pvp-api.minecraft.jp/v1/"

type Session struct {
	client *http.Client
	token  string
}

func New(token string) *Session {
	session := &Session{&http.Client{}, token}
	return session
}

func (session Session) request(endpoint string, query url.Values) ([]byte, error) {

	req, err := http.NewRequest("GET", baseURL+endpoint+"?"+query.Encode(), nil)
	req.Header.Set("Authorization", "Bearer "+session.token)

	if err != nil {
		return nil, err
	}

	dump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%s", dump)

	resp, err := session.client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf(ResponceStatusCode[resp.StatusCode])
	}

	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}
