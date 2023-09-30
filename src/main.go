package main

import (
	"fmt"

	"project.com/attendance/api"
	"project.com/attendance/config"

	"github.com/gin-gonic/gin"
)

func main() {

	config.InitConfig()

	router := gin.Default()

	router.POST("/admin/login", api.AdminLoginHandler)

	router.POST("/faculty/login", api.FacultyLoginHandler)

	router.POST("/admin/add_faculty", api.AddFacultyHandler)

	serverPort := fmt.Sprintf(":%s", config.Config.ServerPort)
	router.Run(serverPort)
}
