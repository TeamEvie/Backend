package images

import (
	"github.com/gominima/minima"
)

func UploadShareX() minima.Handler {
	return func(res *minima.Response, req *minima.Request) {
		res.Send("Hello")
	}
}