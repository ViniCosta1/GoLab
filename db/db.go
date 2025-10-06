package db

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/jackc/pgx/v5/pgxpool"
)

var Ctx = context.Background()

func OpenConn() (*pgxpool.Pool) {
	pool, err := pgxpool.New(Ctx, os.Getenv("DB_CONN"))
    if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	// Verify the connection
    if err := pool.Ping(Ctx); err != nil {
        log.Fatal("Unable to ping database:", err)
    }

	fmt.Println("Connected to PostgreSQL database!")
	return pool
}