package routes

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/gominima/minima"
)

func TestRoute() minima.Handler {
	return func(res *minima.Response, req *minima.Request) {
		res.Status(200)
		param := req.GetParam("id")
		color.Cyan("Got a request from %v", param)
		res.Send(fmt.Sprintf("Hello %v!", param))
	}
}
