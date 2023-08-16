package main

import (
	"goapi/config"
	"goapi/route"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func main() {
	defer config.DisconnectDB(db)
	route.Router()

}
