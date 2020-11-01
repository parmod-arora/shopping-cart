package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	// import pq driver

	"cinemo.com/shoping-cart/framework/appenv"
	_ "github.com/lib/pq"
)

// InitDatabase connect DB
func InitDatabase(databaseURL string) *sql.DB {
	db, err := getDatabase(databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	if i := strings.TrimSpace(os.Getenv("DB_CONN_MAX_OPEN")); i != "" {
		maxOpen, err := strconv.Atoi(i)
		if err == nil {
			db.SetMaxOpenConns(maxOpen)
		}
	}
	if i := strings.TrimSpace(os.Getenv("DB_CONN_MAX_IDLE")); i != "" {
		maxIdle, err := strconv.Atoi(i)
		if err == nil {
			db.SetMaxIdleConns(maxIdle)
		}
	}
	if d := strings.TrimSpace(os.Getenv("DB_CONN_MAX_LIFETIME")); d != "" {
		maxLifetime, err := time.ParseDuration(d)
		if err == nil {
			db.SetConnMaxLifetime(maxLifetime)
		}
	}
	return db
}

func getDatabase(databaseURL string) (*sql.DB, error) {
	return sql.Open("postgres", connectionString(databaseURL))
}

func connectionString(databaseURL string) string {
	if databaseURL != "" {
		if strings.Contains(databaseURL, "sslmode") {
			return databaseURL
		}

		return fmt.Sprintf("%s?sslmode=disable", databaseURL)
	}

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), appenv.GetWithDefault("DB_SCHEMA", "public"))
}
