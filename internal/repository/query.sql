-- name: GetNoteByID :one
SELECT id, title, body FROM notes WHERE id = $1;

-- name: CreateNote :one
INSERT INTO notes (user_id, title, body) VALUES ($1, $2, $3) RETURNING id;

-- name: UpdateNote :exec
UPDATE notes SET title = $2, body = $3 WHERE id = $1;

-- name: DeleteNote :exec
DELETE FROM notes WHERE id = $1;
