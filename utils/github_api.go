package utils

import (
	"encoding/json"
	"time"

	"github.com/fatih/color"
	"github.com/monaco-io/request"
)

type UserRes struct {
	Login             string    `json:"login"`
	ID                int       `json:"id"`
	NodeID            string    `json:"node_id"`
	AvatarURL         string    `json:"avatar_url"`
	GravatarID        string    `json:"gravatar_id"`
	URL               string    `json:"url"`
	HTMLURL           string    `json:"html_url"`
	FollowersURL      string    `json:"followers_url"`
	FollowingURL      string    `json:"following_url"`
	GistsURL          string    `json:"gists_url"`
	StarredURL        string    `json:"starred_url"`
	SubscriptionsURL  string    `json:"subscriptions_url"`
	OrganizationsURL  string    `json:"organizations_url"`
	ReposURL          string    `json:"repos_url"`
	EventsURL         string    `json:"events_url"`
	ReceivedEventsURL string    `json:"received_events_url"`
	Type              string    `json:"type"`
	SiteAdmin         bool      `json:"site_admin"`
	Name              string    `json:"name"`
	Company           string    `json:"company"`
	Blog              string    `json:"blog"`
	Location          string    `json:"location"`
	Email             string    `json:"email"`
	Hireable          bool      `json:"hireable"`
	Bio               string    `json:"bio"`
	TwitterUsername   string    `json:"twitter_username"`
	PublicRepos       int       `json:"public_repos"`
	PublicGists       int       `json:"public_gists"`
	Followers         int       `json:"followers"`
	Following         int       `json:"following"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func GetGHUser(accessToken string) *UserRes {
	var user UserRes
	c := request.Client{
		URL:    "https://api.github.com/user",
		Method: "GET",
		Header: map[string]string{
			"Accept":        "application/json",
			"Authorization": "token " + accessToken,
		},
	}
	resp := c.Send().Do()
	json.Unmarshal(resp.Bytes(), &user)
	return &user
}

func DoesUserSponsorTristan(accessToken string) bool {

	user := GetGHUser(accessToken)

	if user.Login == "twisttaan" {
		color.Cyan("Oh... you're Tristan? I'm impressed!")
		return true
	}

	c := request.Client{
		URL:    "https://api.github.com/graphql",
		Method: "POST",
		Header: map[string]string{
			"Accept":        "application/json",
			"Authorization": "token " + accessToken,
		},
		JSON: map[string]interface{}{
			"query": `
			{
				user(login: "` + user.Login + `") {
				  sponsoring(first: 10) {
					totalCount
					nodes {
					  ... on User {
						login
					  }
					  ... on Organization {
						login
					  }
					}
				  }
				}
			  }
			`,
		},
	}
	resp := c.Send().Do()
	var result struct {
		Data struct {
			User struct {
				Sponsoring struct {
					TotalCount int
					Nodes      []struct {
						Login string
					}
				}
			}
		}
	}

	json.Unmarshal(resp.Bytes(), &result)

	color.HiBlue("%+v", resp.String())

	for _, node := range result.Data.User.Sponsoring.Nodes {
		if node.Login == "twisttaan" {
			return true
		}
	}
	return false
}
