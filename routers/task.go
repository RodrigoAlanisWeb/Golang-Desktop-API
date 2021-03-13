package routers

import (
	"github.com/gin-gonic/gin"
	"rodrigoalanisweb.com/go-desktop-api/controllers"
)

func TaskRoutes(route *gin.Engine) {
	task := route.Group("/task")
	{
		task.GET("/", controllers.RootEndPoint)
		task.POST("/create", controllers.CreateEndPoint)
		task.DELETE("/delete/:id", controllers.DeleteEndPoint)
	}
}
