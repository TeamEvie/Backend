package auth

import (
	"net/http"
	"os"

	"github.com/TeamEvie/Backend/utils"
	"github.com/fatih/color"
	"github.com/gominima/minima"
)

func GitHubAuth() minima.Handler {
	return func(res *minima.Response, req *minima.Request) {

		base := "https://github.com/login/oauth/authorize"
		clientID := os.Getenv("GITHUB_CLIENT_ID")
		scopes := "read:org"

		if req.Raw().URL.Query().Get("code") == "" {
			color.Red("Redirecting to %s", base+"?client_id="+clientID+"&scope="+scopes)
			http.Redirect(res.Raw(), req.Raw(), base+"?client_id="+clientID+"&scope="+scopes, http.StatusTemporaryRedirect)
			return
		}

		code := req.GetQueryParam("code")

		resp := utils.GetAccessToken(code)

		user := utils.GetEvieUserFromGHToken(resp.AccessToken)

		res.Json(user)
	}
}
