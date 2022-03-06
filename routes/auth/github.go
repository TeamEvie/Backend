package auth

import (
	"os"

	"github.com/TeamEvie/Backend/utils"
	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
)

func GitHubAuth(c *fiber.Ctx) error {

	base := "https://github.com/login/oauth/authorize"
	clientID := os.Getenv("GITHUB_CLIENT_ID")
	scopes := "read:org"

	if c.Query("code") == "" {
		color.Red("Redirecting to %s", base+"?client_id="+clientID+"&scope="+scopes+"&redirect_uri="+utils.GetGHRedirectURL())
		return c.Redirect(base + "?client_id=" + clientID + "&scope=" + scopes)
	}

	code := c.Query("code")

	resp := utils.GetAccessToken(code)

	user := utils.GetEvieUserFromGHToken(resp.AccessToken)

	return c.JSON(user)

}
