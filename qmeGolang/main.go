package main

import (
	"QmeServer/controllers"
	"QmeServer/database"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.Connect()

	r.GET("tnh", controllers.GetRecords)
	r.GET("last_tnh", controllers.LastRecord)
	r.POST("tnh", controllers.PostRecord)

	r.GET("led/:id", controllers.GetLed)
	r.POST("led", controllers.PostLed)
	r.PUT("led/:id", controllers.PutLed)

	r.Run("localhost:3000")
}
