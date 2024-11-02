package app

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type App struct {
	DB_INSTANCE *sql.DB
}

func Init() *App {
	return &App{}
}

func (app *App) StartWebServer(ctx context.Context) error {

	envFilePath := "./.env"
	envLoadError := godotenv.Load(envFilePath)

	if envLoadError != nil {
		return fmt.Errorf("failed to load .env: %w", envLoadError)
	}

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	if dbUsername == "" || dbPassword == "" || dbName == "" || dbHost == "" || dbPort == "" {
		return fmt.Errorf("database configuration incomplete. Check environment variables")
	}

	fmt.Printf("DB Username: %s, DB Host: %s, DB Port: %s, DB Name: %s\n",
		dbUsername, dbHost, dbPort, dbName)

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, dbConnectionErr := sql.Open("mysql", dataSourceName)

	if dbConnectionErr != nil {
		return fmt.Errorf("error in connecting db: %w", dbConnectionErr)
	}

	app.DB_INSTANCE = db
	fmt.Printf("connected to the database %s\n", dbName)

	return nil
}
