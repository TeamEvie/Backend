package images

import (
	"context"

	"github.com/TeamEvie/Backend/prisma/db"
	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
)

func UploadShareX(c *fiber.Ctx) error {

	secret := c.GetReqHeaders()["auth"]
	client := db.NewClient()

	if err := client.Prisma.Connect(); err != nil {
		color.Red("[ERROR1] %s", err.Error())
		return c.JSON(fiber.Map{
			"status": "error",
		})
	}

	user, err := client.User.FindFirst(
		db.User.UploadKey.Equals(secret),
	).Exec(context.Background())

	if err != nil {
		color.Red("[ERROR2] %s", err.Error())
		return c.JSON(fiber.Map{
			"status": "Unauthorized",
		})
	}

	if user == nil {
		color.Red("[ERROR3] Unauthorized")
		return c.JSON(fiber.Map{
			"status": "Unauthorized",
		})
	}

	color.Magenta("Secret: \"%s\"", secret)

	return c.JSON(fiber.Map{
		"status": "success",
		"url":    "evie.pw",
	})

}
