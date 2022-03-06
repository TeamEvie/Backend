package auth

import (
	"github.com/TeamEvie/Backend/utils"
	"github.com/gofiber/fiber/v2"
)

func GitHubAuth(c *fiber.Ctx) error {

	if c.Query("code") == "" {
		return c.JSON(fiber.Map{
			"status": "error",
		})
	}

	code := c.Query("code")

	resp := utils.GetAccessToken(code)

	user := utils.GetEvieUserFromGHToken(resp.AccessToken)

	return c.JSON(user)

}
