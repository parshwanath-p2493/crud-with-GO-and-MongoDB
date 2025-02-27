package main

import (
	"FIRST/docs"
	_ "FIRST/docs"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	ConnectDB()
	r := gin.Default()

	swagger := r.Group("swagger")
	{
		docs.SwaggerInfo.Title = "CRUD"
		docs.SwaggerInfo.Description = "CREATE READ UPDATE DELETE"
		docs.SwaggerInfo.Version = "1"

		swagger.GET("/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	//r.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/getdata", GetAll)
	r.POST("/adddata", InsertData)
	r.PUT("/update/:id", UpdateData)
	//r.POST("/as", InsertData)
	r.DELETE("/delete/:id", DeleteData)
	r.GET("/", servercheck)
	r.Run(":5000")

}
