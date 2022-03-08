package images

import (
	"os"

	"github.com/TeamEvie/Backend/prisma/db"
	"github.com/TeamEvie/Backend/utils"
	"github.com/fatih/color"
	"github.com/gominima/minima"
)

func GenSXCU() minima.Handler {
	return func(res *minima.Response, req *minima.Request) {

		secret := req.GetHeader("auth")
		client := db.NewClient()

		if err := client.Prisma.Connect(); err != nil {
			res.InternalServerError().Send(err.Error())
			color.Red("[ERROR1] %s", err.Error())
			return
		}

		user := utils.GetEvieUserFromGHToken(secret)

		if user == nil {
			res.Unauthorized().Send("Unauthorized")
			color.Red("[ERROR2] Unauthorized")
			return
		}

		if user == nil {
			res.Unauthorized().Send("Unauthorized")
			color.Red("[ERROR3] Unauthorized")
			return
		}

		hostname := os.Getenv("BACKEND_URI")
		uploadKey, ok := user.UploadKey()

		if !ok {
			res.InternalServerError().Send("Upload key not found")
			color.Red("[ERROR4] Upload key not found")
			return
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

		res.OK().Send(sxcu)
		color.Green("[SUCCESS] Generated SXCU")
	}
}