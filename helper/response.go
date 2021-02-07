package helper

import "github.com/gin-gonic/gin"

// ResultAPIResponse func
func ResultAPIResponse(str interface{}, length int) gin.H {
	return gin.H{
		"message": "success",
		"data":    str,
		"count":   length,
	}
}

// ResultAPINilResponse func
func ResultAPINilResponse(str interface{}, length int) gin.H {
	return gin.H{
		"message": "success",
		"data":    nil,
		"count":   length,
	}
}
