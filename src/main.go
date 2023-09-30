package main

import (
	"database/sql"
	"fmt"
	"log"

	"project.com/attendance/api"
	"project.com/attendance/config"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {

	connectionString := "user=kmax dbname=test password=2205 sslmode=disable"

	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		return
	}

	// Attempt to ping the database to check the connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
		return
	}

	log.Println("Successfully connected to the PostgreSQL database")
}

func main() {
	config.InitConfig()

	initDB()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	router := gin.Default()

	router.POST("/admin/login", api.AdminLoginHandler)

	router.POST("/faculty/login", api.FacultyLoginHandler)

	router.POST("/admin/add_faculty", api.AddFacultyHandler)

	serverPort := fmt.Sprintf(":%s", config.Config.ServerPort)
	err := router.Run(serverPort)
	if err != nil {
		log.Fatalf("Error starting the server: %v", err)
		return
	}
}
