package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mrwin95/go-ecommerce-backend-api/internal/controllers"
)

func NewRouter() *gin.Engine {

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", controllers.Pong)
		v1.GET("/user/:id", controllers.GetUserById)
	}

	// v2 := r.Group("/api/v2")
	// {
	// 	v2.GET("/ping", Pong)
	// }
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	return r
}
