package helpers

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func DatabaseConfig() (*sql.DB, error) {
	err := godotenv.Load("./config/app.env")
	if err != nil {
		log.Fatalf("เกิดข้อผิดพลาดในการโหลดไฟล์ .env: %v", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", dbUser, dbPass, dbName))

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	return db, err
}
