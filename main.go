package main

import (
	"testcasethree-residentAPI/route"

	"github.com/gin-gonic/gin"
)

func main() {
	app := route.RouteHandler(gin.Default())
	app.Run()
}
