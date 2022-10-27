package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kleawthong/sa-65-test/controller"
	"github.com/kleawthong/sa-65-test/entity"
	
)

const PORT = "3001"

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	router := r.Group("/")
	{
		
			// Admin Routes
			router.GET("/scholar_admins", controller.ListScholarAdmins)
			router.GET("/scholar_admin/:id", controller.GetScholarAdmin)
			router.POST("/scholar_admins", controller.CreateScholarAdmins)
			router.PATCH("/scholar_admins", controller.UpdateScholarAdmin)
			router.DELETE("/scholar_admins/:id", controller.DeleteScholarAdmin)

			// Status Routes
			router.GET("/scholar_statuses", controller.ListScholarStatuses)
			router.GET("/scholar_status/:id", controller.GetScholarStatus)
			router.POST("/scholar_statuses", controller.CreateScholarStatus)
			router.PATCH("/scholar_statuses", controller.UpdateScholarStatus)
			router.DELETE("/scholar_statuses/:id", controller.DeleteScholarStatus)

			// Type Routes
			router.GET("/scholar_types", controller.ListScholarTypes)
			router.GET("/scholar_type/:id", controller.GetScholarType)
			router.POST("/scholar_types", controller.CreateScholarType)
			router.PATCH("/scholar_types", controller.UpdateScholarType)
			router.DELETE("/scholar_types/:id", controller.DeleteScholarType)

			// Scholarship Routes
			router.GET("/scholarships", controller.ListScholarships)
			router.GET("/scholarship/:id", controller.GetScholarship)
			router.POST("/scholarships", controller.CreateScholarship)
			router.PATCH("/scholarships", controller.UpdateScholarship)
			router.DELETE("/scholarships/:id", controller.DeleteScholarship)

		
	}
	r.Run("localhost: " + PORT)
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
