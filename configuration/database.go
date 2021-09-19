package configuration

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {
	connection := "andri:root@tcp(127.0.0.1:3306)/order_by?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", connection)

	if err != nil {
		fmt.Println(err)
		panic("failed connect to database")
	}

	return db
}
