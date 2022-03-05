package utils

import (
	"encoding/json"
	"os"

	"github.com/monaco-io/request"
)

type AccessTokenRes struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

func GetAccessToken(code string) *AccessTokenRes {
	clientID := os.Getenv("GITHUB_CLIENT_ID")
	clientSecret := os.Getenv("GITHUB_CLIENT_SECRET")

	var result AccessTokenRes

	c := request.Client{
		URL:    "https://github.com/login/oauth/access_token",
		Method: "POST",
		Header: map[string]string{
			"Accept": "application/json",
		},
		JSON: map[string]string{
			"client_id":     clientID,
			"client_secret": clientSecret,
			"code":          code,
		},
	}

	resp := c.Send().Do()

	json.Unmarshal(resp.Bytes(), &result)

	return &result
}
