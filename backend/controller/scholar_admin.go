package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kleawthong/sa-65-test/entity"
)
func CreateScholarAdmins(c *gin.Context) {
	var scholar_admins entity.ScholarStatus
	if err := c.ShouldBindJSON(&scholar_admins); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&scholar_admins).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": scholar_admins})

}
func ListScholarAdmins(c *gin.Context) {
	var scholar_admins []entity.ScholarAdmin
	if err := entity.DB().Raw("SELECT * FROM scholar_admins").Scan(&scholar_admins).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": scholar_admins})
}

// GET /scholar_admin/:id
// Get scholar_admin by id
func GetScholarAdmin(c *gin.Context) {
	var scholar_admin entity.ScholarAdmin
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&scholar_admin); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "scholar_admins not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": scholar_admin})
}

// PATCH /scholar_admins
func UpdateScholarAdmin(c *gin.Context) {
	var scholar_admin entity.ScholarAdmin
	if err := c.ShouldBindJSON(&scholar_admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", scholar_admin.ID).First(&scholar_admin); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "scholar_admins not found"})
		return
	}

	if err := entity.DB().Save(&scholar_admin).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": scholar_admin})
}

// DELETE /scholar_adminss/:id
func DeleteScholarAdmin(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM scholar_admins WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	/*
		if err := entity.DB().Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}*/

	c.JSON(http.StatusOK, gin.H{"data": id})
}