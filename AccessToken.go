package jpmcpvpapi

type ClientCredentials struct {
	AccessToken string      `json:"access_token"`
	ExpiresIn   int         `json:"expires_in"`
	TokenType   string      `json:"token_type"`
	Scope       interface{} `json:"scope"`
}

func GetAccessToken(client_id string, client_secret string) (client *ClientCredentials, err error) {

}
