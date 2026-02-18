package main

import (
	"context"
	"log"
	"net/http"
	"notes/internal/repository/db"
	"notes/internal/router"

	"github.com/jackc/pgx/v5"
)

func main() {
	connStr := "postgres://user:password@localhost:5433/notes_db?sslmode=disable"
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())
	queries := db.New(conn)
	r := router.NewRouter(queries)
	log.Println("Starting server on :3333")
	http.ListenAndServe(":3333", r)
}
