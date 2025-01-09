package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/Vkanhan/code-bin/internal/models"
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
		log.Fatal("Error loading .env file")
	}

	addr := os.Getenv("PORT")
	if addr == "" {
		addr = ":4000" // Default to ":4000" if not set
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := connectToDB()
	if err != nil {
		log.Fatal(err)
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

	infoLog.Printf("listening to server on port: %s", addr)
	err = server.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}

func connectToDB() (*sql.DB, error) {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Print("Database connection error")
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	return db, err
}
