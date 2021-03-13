package controllers

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex"`
	Password string
}

type JwtToken struct {
	Token string `json:"token"`
}

func RegisterEndPoint(c *gin.Context) {
	connect()
	username := c.PostForm("username")
	password := c.PostForm("password")
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": "Error",
		})
		return
	}
	user := User{Username: username, Password: string(hash)}
	res := db.Create(&user)
	if res.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": "Error",
		})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": "Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func VerifyToken(c *gin.Context) int {
	token := c.Request.Header.Get("x-access-token")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": "You Dont Provide A Token",
		})
		return 0
	}

	vtoken, _ := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret"), nil
	})

	if claims, ok := vtoken.Claims.(jwt.MapClaims); ok && vtoken.Valid {
		var userId int
		mapstructure.Decode(claims["id"], &userId)
		return userId
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg": "Invalid Token",
		})
		return 0
	}
}
