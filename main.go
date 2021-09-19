package main

import (
	"assigment2-glng-restapi/configuration"
	"assigment2-glng-restapi/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := configuration.InitDB()

	DBConn := &controllers.DBConn{db}
	router := gin.Default()

	router.GET("/orders", DBConn.GetOrders)
	router.GET("/order/:id", DBConn.GetOrder)
	router.DELETE("/order/:id", DBConn.DeleteOrder)
	router.POST("/order", DBConn.CreateOrder)
	router.PUT("/order", DBConn.UpdatOrder)

	router.Run(":8000")
}
