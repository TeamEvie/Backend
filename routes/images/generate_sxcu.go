package images

import (
	"os"

	"github.com/TeamEvie/Backend/prisma/db"
	"github.com/TeamEvie/Backend/utils"
	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
)

func GenSXCU(c *fiber.Ctx) error {

	secret := c.GetReqHeaders()["auth"]
	client := db.NewClient()

	if err := client.Prisma.Connect(); err != nil {
		color.Red("[ERROR1] %s", err.Error())
		return c.JSON(fiber.Map{
			"status": "error",
		})
	}

	user := utils.GetEvieUserFromGHToken(secret)

	if user == nil {
		color.Red("[ERROR2] Unauthorized")
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

	hostname := os.Getenv("BACKEND_URI")
	uploadKey, ok := user.UploadKey()

	if !ok {
		color.Red("[ERROR4] Upload key not found")
		return c.JSON(fiber.Map{
			"status": "Upload key not found",
		})
	}

	sxcu := `{
			"Version": "13.7.0",
			"Name": "Local",
			"DestinationType": "ImageUploader, FileUploader",
			"RequestMethod": "POST",
			"RequestURL": "` + hostname + `/v1/sharex",
			"Headers": {
			  "auth": "` + uploadKey + `"
			},
			"Body": "Binary",
			"URL": "$json:url$"
		  }`

	return c.SendFile(sxcu)

}
