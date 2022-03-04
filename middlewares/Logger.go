package middlewares

import (
	"github.com/fatih/color"
	"github.com/gominima/minima"
)

func Logger(res *minima.Response, req *minima.Request) {
	color.Magenta("Request incoming for \"%s\"", req.GetPathURL())
}
