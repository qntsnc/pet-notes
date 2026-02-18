package router

import (
	"notes/internal/handlers"
	"notes/internal/repository/db"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(repo *db.Queries) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	noteHandler := &handlers.NoteHandler{Repo: repo}
	r.Get("/", handlers.HomeHandler)
	r.Post("/notes", noteHandler.PostNote)
	r.Get("/notes/{id}", noteHandler.GetNoteByID)
	return r
}
