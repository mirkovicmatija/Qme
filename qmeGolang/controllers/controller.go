package controllers

import (
	"QmeServer/database"
	"QmeServer/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type SensorDataInput struct {
	Temperature float32   `json:"temperature" binding:"required"`
	Humidity    float32   `json:"humidity" binding:"required"`
	Time        time.Time `json:"timestamp"`
}

type LedDataInput struct {
	State int `json:"state"`
}

func GetRecords(c *gin.Context) {
	var records []model.SensorData
	database.DB.Find(&records)
	c.JSON(http.StatusOK, gin.H{"data": records})
}

func LastRecord(c *gin.Context) {
	var lastRecord model.SensorData
	database.DB.Last(&lastRecord)
	c.JSON(http.StatusOK, gin.H{"data": lastRecord})
}

func PostRecord(c *gin.Context) {
	var input SensorDataInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data := model.SensorData{Timestamp: time.Now(), Temperature: input.Temperature, Humidity: input.Humidity}
	database.DB.Create(&data)

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func GetLed(c *gin.Context) {
	var led model.LedData
	ledId := c.Param("id")
	database.DB.First(&led, ledId)
	c.JSON(http.StatusOK, gin.H{"data": led})

}

func PostLed(c *gin.Context) {
	data := model.LedData{State: 0}
	database.DB.Create(&data)
	c.JSON(http.StatusOK, gin.H{"data": data})

}

func PutLed(c *gin.Context) {
	var led model.LedData
	ledId := c.Param("id")
	database.DB.First(&led, ledId)
	if led.State == 1 {
		database.DB.Model(&led).Update("state", 0)
	} else {
		database.DB.Model(&led).Update("state", 1)
	}
	c.JSON(http.StatusOK, gin.H{"data": led})
}
