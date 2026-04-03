-- name: GetTodos :many
SELECT * FROM todos ORDER BY id DESC;

-- name: GetTodo :one
SELECT * FROM todos WHERE id=$1 LIMIT 1;

-- name: UpdateTodoComplete :one
UPDATE todos
SET complete = $2
WHERE id = $1
RETURNING *;

-- name: InsertTodo :one
INSERT INTO todos (title) VALUES ($1) RETURNING *;
