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
			res.Status(500).Send(err.Error())
			color.Red("[ERROR1] %s", err.Error())
			return
		}

		user, err := client.User.FindFirst(
			db.User.UploadKey.Equals(secret),
		).Exec(context.Background())

		if err != nil {
			res.Status(500).Send(err.Error())
			color.Red("[ERROR2] %s", err.Error())
			return
		}

		if user == nil {
			res.Status(401).Send("Unauthorized")
			color.Red("[ERROR3] Unauthorized")
			return
		}

		color.Magenta("Secret: \"%s\"", secret)

		res.Json(map[string]string{"url": "test.com"})
	}
}
