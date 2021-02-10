package helper

import "github.com/jasonlvhit/gocron"

// RunCron func
func RunCron() {
	gocron.Every(5).Seconds().Do(task)
	<-gocron.Start()
}
