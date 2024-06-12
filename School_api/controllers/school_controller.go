package controllers

import (
	"net/http"
	"strconv"

	"School_api/models"
	"School_api/repositories"
	"School_api/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type SchoolController struct {
	service services.SchoolService
}

func NewSchoolController(db *gorm.DB) *SchoolController {
	repo := repositories.NewSchoolRepository(db)
	service := services.NewSchoolService(repo)
	return &SchoolController{service}
}

func (ctrl *SchoolController) GetSchools(c *gin.Context) {
	schools, err := ctrl.service.GetSchools()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, schools)
}

func (ctrl *SchoolController) GetSchoolByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	school, err := ctrl.service.GetSchoolByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "School not found"})
		return
	}
	c.JSON(http.StatusOK, school)
}

func (ctrl *SchoolController) CreateSchool(c *gin.Context) {
	var school models.School
	if err := c.ShouldBindJSON(&school); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdSchool, err := ctrl.service.CreateSchool(school)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, createdSchool)
}

func (ctrl *SchoolController) UpdateSchool(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var school models.School
	if err := c.ShouldBindJSON(&school); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedSchool, err := ctrl.service.UpdateSchool(uint(id), school)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedSchool)
}

func (ctrl *SchoolController) DeleteSchool(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.service.DeleteSchool(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "School deleted"})
}
