package router

import (
	"go_harians/config"
	"gorm.io/gorm"
)

var db *gorm.DB = config.COnnectionDB()
