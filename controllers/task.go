package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db gorm.DB

type Task struct {
	gorm.Model
	ID     uint `gorm:"primaryKey"`
	Name   string
	Done   bool
	UserId int
}

func connect() {
	dsn := "root@tcp(127.0.0.1:3306)/desktop-api"
	con, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db = *con
}

func RootEndPoint(c *gin.Context) {
	userId := VerifyToken(c)
	connect()

	var result []Task
	db.Raw("SELECT id,name,done,user_id,created_at,updated_at FROM tasks WHERE user_id = ?", userId).Scan((&result))

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func DeleteEndPoint(c *gin.Context) {
	userId := VerifyToken(c)
	connect()

	db.Exec("DELETE FROM tasks WHERE user_id = ? AND id = ?", userId, c.Param("id"))

	c.JSON(http.StatusOK, gin.H{
		"msg": "Deleted!",
	})
}

func CreateEndPoint(c *gin.Context) {
	userId := VerifyToken(c)
	connect()
	name := c.PostForm("name")
	task := Task{Name: name, Done: false, UserId: userId}
	res := db.Create(&task)
	if res.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": "Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Created",
		"id":  task.ID,
	})
}
