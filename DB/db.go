package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Could not initialize database connection: %v", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatalf("Database connection is not available: %v", err)
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(5 * time.Minute)
	log.Println("Database connection successful.")
	createTables()
}

func createTables() {

	createAIRequestsTable := `
    CREATE TABLE IF NOT EXISTS ai_requests (
        id INT PRIMARY KEY AUTO_INCREMENT,
        original_text TEXT NOT NULL,
        summary_text TEXT,
        prompt TEXT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`

	_, err := DB.Exec(createAIRequestsTable)
	if err != nil {
		log.Fatalf("Failed to create ai_requests table: %v", err)
	}
	log.Println("ai_requests table checked/created successfully.")
}
