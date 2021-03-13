package routers

import (
	"github.com/gin-gonic/gin"
	"rodrigoalanisweb.com/go-desktop-api/controllers"
)

func AuthRoutes(route *gin.Engine) {
	task := route.Group("/task")
	{
		task.POST("/register", controllers.RegisterEndPoint)
	}
}
