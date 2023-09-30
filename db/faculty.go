package db

import (
	"database/sql"
	"fmt"
	"log"
	"project.com/attendance/models"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	dataSourceName := fmt.Sprintf("user=kmax password=2205 dbname=test host=localhost port=5432 sslmode=disable")

	var err error

	db, err = sql.Open(`postgres`, dataSourceName)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	createFacultyTable()
}

func createFacultyTable() {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS faculty (
			id SERIAL PRIMARY KEY,
			username TEXT NOT NULL,
			password TEXT NOT NULL
		)
	`)
	if err != nil {
		log.Fatalf("Error creating faculty table: %v", err)
	}
}

func GetFacultyByUsernameAndPassword(username, password string) (*models.Faculty, error) {
	var faculty models.Faculty
	err := db.QueryRow("SELECT id, username, password FROM faculty WHERE username = $1 AND password = $2", username, password).Scan(&faculty.ID, &faculty.Username, &faculty.Password)
	if err != nil {
		return nil, err
	}
	return &faculty, nil
}

func UsernameExists(username string) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM faculty WHERE username = $1", username).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func AddFaculty(faculty *models.Faculty) (int, error) {
	var lastInsertID int
	err := db.QueryRow("INSERT INTO faculty (username, password) VALUES ($1, $2) RETURNING id", faculty.Username, faculty.Password).Scan(&lastInsertID)
	if err != nil {
		log.Printf("Error inserting faculty: %v", err)
		return 0, err
	}
	return lastInsertID, nil
}
