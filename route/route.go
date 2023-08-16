package route

import (
	"goapi/controllers"
	_ "goapi/controllers"

	"github.com/gin-gonic/gin"
)

func Router() {

	route := gin.Default()

	api := route.Group("/api/v1")
	{
		user := api.Group("/user")
		user.GET("/", controllers.ListUserAll)
		//user.GET(":userid", controllers.ListUserByid)
		user.POST("/", controllers.CreateUser)
		user.PUT(":userid", controllers.UpdateUser)
		user.DELETE(":userid", controllers.DeteleUser)

		product := api.Group("/products")
		product.GET("/", controllers.ListProductAll)
		product.POST(":userid", controllers.ListProductByid)
		product.POST("/", controllers.CreateProduct)
		product.PUT(":userid", controllers.UpdateProduct)
		product.DELETE(":userid", controllers.DeteleProduct)

	}
	route.Run("localhost:8081")
}
