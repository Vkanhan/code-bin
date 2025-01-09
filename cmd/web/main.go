package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Vkanhan/code-bin/internal/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	gists    *models.GistModel
	users    *models.UserModel
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("error loading .env file")
	}

	addr := os.Getenv("PORT")
	if addr == "" {
		log.Println("PORT is not available in .env file")
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := connectToDB()
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		gists:    &models.GistModel{DB: db},
		users:    &models.UserModel{DB: db},
	}

	server := &http.Server{
		Addr:     addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", addr)
	err = server.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}

func connectToDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("error loading .env")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	// Add connection retry logic
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
		if i < maxRetries-1 {
			time.Sleep(time.Second * 5) // Wait 5 seconds before retrying
		}
	}
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database after %d attempts: %v", maxRetries, err)
	}

	return db, nil
}
