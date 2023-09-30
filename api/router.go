package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/admin/login", AdminLoginHandler)
	router.POST("/faculty/login", FacultyLoginHandler)
	router.POST("/admin/add_faculty", AddFacultyHandler)
	return router
}

func StartServer() {
	router := SetupRouter()
	router.Run(":8083")
}
