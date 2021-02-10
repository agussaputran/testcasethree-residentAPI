package main

import (
	"testcasethree-residentAPI/others"
	"testcasethree-residentAPI/route"

	"github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"
)

// RunCron func
func RunCron() {
	gocron.Every(10).Seconds().Do(others.CheckQueue)
	<-gocron.Start()
}

func main() {
	go RunCron()
	app := route.RouteHandler(gin.Default())
	app.Run()
}
