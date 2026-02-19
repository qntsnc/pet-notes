// cmd/server/server.go

package main

import (
	"context"
	"log"
	"net/http"
	"time" // Импортируем пакет time

	"notes/internal/repository/db"
	"notes/internal/router"

	"github.com/jackc/pgx/v5" // Используем pgxpool для лучшего управления соединениями
)

func main() {
	connStr := "postgres://server:password@postgres:5432/notes_db?sslmode=disable"

	var conn *pgx.Conn
	var err error

	// Цикл для повторных попыток подключения
	for i := 0; i < 5; i++ {
		conn, err = pgx.Connect(context.Background(), connStr)
		if err == nil {
			log.Println("Successfully connected to the database!")
			break // Выходим из цикла, если подключение успешно
		}

		log.Printf("Failed to connect to database (attempt %d/5): %v\n", i+1, err)
		log.Println("Retrying in 2 seconds...")
		time.Sleep(2 * time.Second)
	}

	// Если после всех попыток подключиться не удалось, завершаем работу
	if err != nil {
		log.Fatalf("Unable to connect to database after several attempts: %v\n", err)
	}

	queries := db.New(conn)
	r := router.NewRouter(queries)

	log.Println("Starting server on :3333")
	http.ListenAndServe(":3333", r)
}
