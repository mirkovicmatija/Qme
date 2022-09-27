package model

import "time"

type SensorData struct {
	Id          uint      `json:"id" gorm:"primary_key"`
	Timestamp   time.Time `json:"timestamp"`
	Temperature float32   `json:"temperature" sql:"type:decimal"`
	Humidity    float32   `json:"humidity" sql:"type:decimal"`
}

type LedData struct {
	Id    uint `json:"id" gorm:"primary_key"`
	State uint `json:"state"`
}
