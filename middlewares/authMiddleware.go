package middlewares

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testcasethree-residentAPI/connection"
	"testcasethree-residentAPI/helper"
	"testcasethree-residentAPI/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Auth func
func Auth(c *gin.Context) {
	secret := helper.GetEnvVar("JWT_SECRET")
	tokenStringHeader := c.Request.Header.Get("Authorization")
	allowedMethod := c.Request.Method
	reqPath := c.Request.URL.Path
	token, err := jwt.Parse(tokenStringHeader, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Method tidak diketahui atau bukan HS256 %V", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if token != nil && err == nil {
		payload := token.Claims.(jwt.MapClaims)
		log.Println("Token Verified")

		if payload["role"] == "entry" {
			var queue models.QueueEmail
			db := connection.Connect()
			buf, _ := ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
			queue.To = "admin@yopmail.com"
			queue.Cc = ""
			queue.Subject = string(allowedMethod) + string(reqPath)
			queue.Handled = false
			queue.Message = string(buf)
			db.Create(&queue)
		}

		if payload["role"] == "guest" && allowedMethod != "GET" {
			result := gin.H{
				"message": "You can't access this route",
			}
			LogSentryUserRequest(payload, c)
			c.Abort()
			c.JSON(http.StatusUnauthorized, result)
		} else if payload["role"] == "entry" && allowedMethod == "DELETE" {
			result := gin.H{
				"message": "You can't access this route",
			}
			LogSentryUserRequest(payload, c)
			c.Abort()
			c.JSON(http.StatusUnauthorized, result)
		} else {
			LogTerminalUserRequest(payload, c)
		}

	} else if err != nil {
		log.Println("Wrong Token, error -> ", err.Error())
		result := gin.H{
			"message": "Token is not valid",
			"error":   http.StatusUnauthorized,
		}
		LogTerminalRequest(c)
		c.Abort()
		c.JSON(http.StatusUnauthorized, result)
	}
}
