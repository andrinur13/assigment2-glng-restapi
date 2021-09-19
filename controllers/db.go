package controllers

import "github.com/jinzhu/gorm"

type DBConn struct {
	DB *gorm.DB
}
