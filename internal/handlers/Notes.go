package handlers

import (
	"encoding/json"
	"net/http"
	"notes/internal/repository/db"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type NoteHandler struct {
	Repo *db.Queries
}

func (h *NoteHandler) PostNote(w http.ResponseWriter, r *http.Request) {
	var params db.CreateNoteParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	noteID, err := h.Repo.CreateNote(r.Context(), params)
	if err != nil {
		http.Error(w, "Failed to create note", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Note created successfully",
		"note_id": noteID,
	})
}

func (h *NoteHandler) GetNoteByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}
	note, err := h.Repo.GetNoteByID(r.Context(), id)

	if err != nil {
		http.Error(w, "Failed to get note", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(note)
}

func (h *NoteHandler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	var params db.UpdateNoteParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Repo.UpdateNote(r.Context(), params); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Note updated successfully",
	})
}
