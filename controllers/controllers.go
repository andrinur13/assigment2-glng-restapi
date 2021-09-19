package controllers

import (
	"assigment2-glng-restapi/structs"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (Conn *DBConn) GetOrders(c *gin.Context) {
	var ordersData []structs.Order

	Conn.DB.Find(&ordersData)

	for i := range ordersData {
		var items []structs.Item
		Conn.DB.Where("order_id = ?", ordersData[i].OrderId).Find(&items)
		ordersData[i].Items = items
	}

	fmt.Println(ordersData)

	if len(ordersData) < 1 {
		c.JSON(http.StatusOK, gin.H{
			"result": "data not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": ordersData,
	})
	return
}

func (Conn *DBConn) GetOrder(c *gin.Context) {
	var orderData structs.Order
	var items []structs.Item

	var id string = c.Param("id")
	id_order, err := strconv.Atoi(id)

	if err != nil {
		log.Fatal(err)
	}

	Conn.DB.Where("order_id = ?", id_order).Find(&orderData)
	Conn.DB.Where("order_id = ?", id_order).Find(&items)

	orderData.Items = items

	if orderData.OrderId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"result": "Data not found!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": orderData,
	})
	return
}

func (Conn *DBConn) DeleteOrder(c *gin.Context) {
	var orderData structs.Order
	var items []structs.Item

	var id string = c.Param("id")
	id_order, err := strconv.Atoi(id)

	if err != nil {
		log.Fatal(err)
	}

	Conn.DB.Where("order_id = ?", id_order).Find(&orderData)

	if orderData.OrderId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"result": "Data not found!",
		})
		return
	}

	Conn.DB.Where("order_id = ?", id_order).Delete(&orderData)

	Conn.DB.Where("order_id = ?", id_order).Find(&items)
	Conn.DB.Delete(&items)

	c.JSON(http.StatusOK, gin.H{
		"result": "success deleted data",
	})
	return

}

func (Conn *DBConn) CreateOrder(c *gin.Context) {
	var orderData structs.Order
	var items []structs.Item

	err := c.ShouldBindJSON(&orderData)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result": "Terjadi kesalahan saat membaca json",
			"error":  err,
		})
		return
	}

	items = orderData.Items

	// create data orders dulu
	Conn.DB.Create(&orderData)

	for i := range items {
		items[i].OrderId = orderData.OrderId
		Conn.DB.Create(items[i])
	}

	c.JSON(http.StatusOK, gin.H{
		"result": orderData,
	})
}

func (Conn *DBConn) UpdatOrder(c *gin.Context) {
	var orderData structs.Order
	var items []structs.Item

	err := c.ShouldBindJSON(&orderData)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result": "Terjadi kesalahan saat membaca json",
			"error":  err,
		})
		return
	}

	items = orderData.Items

	// update
	updatedOrder := structs.Order{}
	Conn.DB.Where("order_id = ?", orderData.OrderId).First(&updatedOrder)
	updatedOrder.CustomerName = orderData.CustomerName
	updatedOrder.OrderedAt = orderData.OrderedAt
	Conn.DB.Save(&updatedOrder)

	for i := range items {
		var updatedItem structs.Item

		Conn.DB.Where("item_id = ?", items[i].ItemId).First(&updatedItem)
		updatedItem.ItemCode = items[i].ItemCode
		updatedItem.Description = items[i].Description
		updatedItem.Quantity = items[i].Quantity
		Conn.DB.Save(&updatedItem)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": orderData,
	})
}
