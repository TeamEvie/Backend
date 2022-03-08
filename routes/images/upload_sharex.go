package images

import (
	"context"
	"github.com/TeamEvie/Backend/prisma/db"
	"github.com/fatih/color"
	"github.com/gominima/minima"
)

func UploadShareX() minima.Handler {
	return func(res *minima.Response, req *minima.Request) {
		secret := req.GetHeader("auth")
		client := db.NewClient()

		if err := client.Prisma.Connect(); err != nil {
			defer client.Prisma.Disconnect()
			color.Red("[ERROR1] %s", err.Error())
			res.InternalServerError().Send(err.Error())
			return
		}

		user, err := client.User.FindFirst(
			db.User.UploadKey.Equals(secret),
		).Exec(context.Background())

		if err != nil {
			color.Red("[ERROR2] %s", err.Error())
			res.InternalServerError().Send(err.Error())
			return
		}

		if user == nil {
			color.Red("[ERROR3] Unauthorized")
			res.Unauthorized().Send("Unauthorized")
			return
		}

		color.Magenta("Secret: \"%s\"", secret)

		res.JSON(map[string]string{"status": "success", "url": "evie.pw"})
	}
}
