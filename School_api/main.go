package main

import (
	"School_api/controllers"
	"School_api/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize the database
	db := utils.InitDB()

	// Inject the database into the controllers
	schoolController := controllers.NewSchoolController(db)

	// School routes
	r.GET("/schools", schoolController.GetSchools)
	r.GET("/schools/:id", schoolController.GetSchoolByID)
	r.POST("/schools", schoolController.CreateSchool)
	r.PUT("/schools/:id", schoolController.UpdateSchool)
	r.DELETE("/schools/:id", schoolController.DeleteSchool)

	r.Run(":8080")
}
