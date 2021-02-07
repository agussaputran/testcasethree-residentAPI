package controllers

import (
	"log"
	"net/http"
	"testcasethree-residentAPI/helper"
	"testcasethree-residentAPI/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LoginUser func
func (gorm *Gorm) LoginUser(c *gin.Context) {
	var (
		user   models.Users
		userDB models.Users
		result gin.H
	)

	if err := c.Bind(&user); err != nil {
		log.Println("Data tidak ada, error : ", err.Error())
	}

	gorm.DB.Where("email = ?", user.Email).First(&userDB)

	if err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password)); err != nil {
		log.Println("Email ", user.Email, " Password salah")
		result = gin.H{
			"message": "email atau password salah",
		}
	} else {
		type authCustomClaims struct {
			Email string `json:"email"`
			Role  string `json:"role"`
			jwt.StandardClaims
		}

		claims := &authCustomClaims{
			userDB.Email,
			userDB.Role,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
		}
		sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
		token, err := sign.SignedString([]byte(helper.GetEnvVar("JWT_SECRET")))
		if err != nil {
			log.Println("Gagal create token, message ", err.Error())
			result = gin.H{
				"message": "Gagal create token",
				"token":   nil,
			}
		} else {
			log.Println("Email ", user.Email, " Berhasil login")
			result = gin.H{
				"message": "anda berhasil login",
				"token":   token,
			}
		}
	}

	c.JSON(http.StatusOK, result)
}
