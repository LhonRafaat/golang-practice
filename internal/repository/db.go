package repository

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)


var DB *pgxpool.Pool


func InitDB() {
    var err error
    connStr := os.Getenv("DATABASE_URL")
    DB, err = pgxpool.New(context.Background(), connStr)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
    }
}

func CloseDB() {
    DB.Close()
}