package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	host := getEnv("DB_HOST", "localhost")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "postgres")
	dbname := getEnv("DB_NAME", "finance_db")
	port := getEnv("DB_PORT", "5432")

	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)

	var db *sql.DB
	var err error

	for i := 1; i <= 5; i++ {
		log.Printf("Connecting to DB... attempt %d", i)

		db, err = sql.Open("postgres", connStr)
		if err == nil {
			err = db.Ping()
			if err == nil {
				break
			}
		}

		log.Println("DB not ready, retrying in 2 seconds...")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal(" Could not connect to DB:", err)
	}

	DB = db
	log.Println(" Connected to PostgreSQL")
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
