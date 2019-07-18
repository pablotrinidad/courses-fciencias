// Setup and initialize DB

package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func init() {
	dbName := getenv("POSTGRES_DB", "")
	dbUser := getenv("POSTGRES_USER", "")
	dbPassword := getenv("POSTGRES_PASSWORD", "")
	dbHost := getenv("POSTGRES_HOST", "fciencias_db")
	dbPort := getenv("POSTGRES_PORT", "5432")
	dbSSLMode := getenv("POSTGRES_SSL_MODE", "disable")

	connection := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbName, dbPassword, dbSSLMode)

	db, err = gorm.Open("postgres", connection)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	runMigrations()
}

// getenv retrieve an environment value and return a default value if it doesnt exist.
func getenv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

// runMigrations goes through model definitions and apply GORM AutoMigrations
func runMigrations() {
	// Majors
	db.AutoMigrate(&Major{})
	db.AutoMigrate(&AcademicProgram{})
}

// GetDB return a pointer to an initialized GORM DB client
func GetDB() *gorm.DB {
	return db
}
