package api

import (
	"net/http"

	"project.com/attendance/db"
	"project.com/attendance/models"

	"github.com/gin-gonic/gin"
)

func AdminLoginHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "test" && password == "123" {
		c.JSON(http.StatusOK, gin.H{"message": "Admin login successful"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}

func FacultyLoginHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	faculty, err := db.GetFacultyByUsernameAndPassword(username, password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, faculty)
}

func AddFacultyHandler(c *gin.Context) {
	var faculty models.Faculty
	if err := c.ShouldBindJSON(&faculty); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	exists, err := db.UsernameExists(faculty.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error checking username"})
		return
	}

	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	lastInsertID, err := db.AddFaculty(&faculty)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error adding faculty"})
		return
	}

	faculty.ID = lastInsertID

	c.JSON(http.StatusCreated, faculty)
}
