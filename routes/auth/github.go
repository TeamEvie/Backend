package auth

import (
	"github.com/TeamEvie/Backend/utils"
	"github.com/gominima/minima"
)

func GitHubAuth() minima.Handler {
	return func(res *minima.Response, req *minima.Request) {
		if req.Query("code") == "" {
			res.InternalServerError().JSON(map[string]string{"status": "error"})
			return
		}

		code := req.Query("code")

		resp := utils.GetAccessToken(code)

		// user := utils.GetEvieUserFromGHToken(resp.AccessToken)

		res.OK().JSON(map[string]string{
			"status": "success",
			"token":  resp.AccessToken,
		})
	}
}
