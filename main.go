package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"rodrigoalanisweb.com/go-desktop-api/controllers"
	"rodrigoalanisweb.com/go-desktop-api/routers"
)

func main() {
	dsn := "root@tcp(127.0.0.1:3306)/desktop-api"
	con, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	con.AutoMigrate(&controllers.Task{})
	con.AutoMigrate(&controllers.User{})

	router := gin.Default()

	routers.AuthRoutes(router)
	routers.TaskRoutes(router)

	router.Run()
}
