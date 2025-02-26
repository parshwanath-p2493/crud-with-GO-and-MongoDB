package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	ConnectDB()
	r := gin.Default()
	r.GET("/getdata", GetAll)
	r.POST("/adddata", InsertData)
	r.PUT("/update/:id", UpdateData)
	//r.POST("/as", InsertData)
	r.DELETE("/delete/:id", DeleteData)
	r.GET("/", servercheck)
	r.Run(":5000")

}
