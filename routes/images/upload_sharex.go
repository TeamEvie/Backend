package images

import (
	"github.com/fatih/color"
	"github.com/gominima/minima"
)

func UploadShareX() minima.Handler {
	return func(res *minima.Response, req *minima.Request) {
		// the incoming request is a multipart/form-data request get the secret and file

		secret := req.GetHeader("auth")

		color.Magenta("Secret: \"%s\"", secret)

		res.Json(map[string]string{"url": "test.com"})
	}
}
