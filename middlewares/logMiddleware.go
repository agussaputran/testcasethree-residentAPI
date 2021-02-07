package middlewares

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

// LogTerminalUserRequest for logging user request to terminal
func LogTerminalUserRequest(payload map[string]interface{}, c *gin.Context) {
	reqMethod := c.Request.Method
	reqPath := c.Request.URL.Path
	email := fmt.Sprintf("%v", payload["email"])
	role := fmt.Sprintf("%v", payload["role"])
	logMessage := "email : " + string(email) + " || role : " + role + " || action : " + string(reqMethod) + " -> " + reqPath
	log.Println(logMessage)
}

// LogSentryUserRequest func for logging user request to sentry.io
func LogSentryUserRequest(payload map[string]interface{}, c *gin.Context) {
	reqMethod := c.Request.Method
	reqPath := c.Request.URL.Path
	buf, _ := ioutil.ReadAll(c.Request.Body)
	email := fmt.Sprintf("%v", payload["email"])
	role := fmt.Sprintf("%v", payload["role"])
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
	sentryMessage := "email : " + string(email) + " || role : " + role + " || action : " + string(reqMethod) + " -> " + reqPath + "\n" + string(buf)
	Sentry(sentryMessage)
}

// Sentry func to log with sentry.io
func Sentry(data string) {
	dsn := os.Getenv("DSN")

	err := sentry.Init(sentry.ClientOptions{
		// Either set your DSN here or set the SENTRY_DSN environment variable.
		Dsn: dsn,
	})
	if err != nil {
		sentry.CaptureException(err)
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	defer sentry.Flush(2 * time.Second)
	sentry.CaptureMessage(data)
}
