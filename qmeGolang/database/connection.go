package database

import (
	"QmeServer/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, error := gorm.Open(mysql.Open("root:Grof2003#@/qme?parseTime=true"), &gorm.Config{})
	if error != nil {
		panic("Could not connect do db")
	}
	DB = connection

	connection.AutoMigrate(&model.SensorData{}, &model.LedData{})
}
