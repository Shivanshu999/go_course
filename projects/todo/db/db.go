package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func ConnectDB() (*pgx.Conn, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Error loading .env file: %v", err)
	}

	databaseUrl := os.Getenv("DATABASE_URL")

	if databaseUrl == "" {
		return nil, fmt.Errorf("database url not set")

	}
	conn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}
	return conn, nil
}
