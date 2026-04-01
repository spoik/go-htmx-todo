-- name: GetTodos :many
SELECT * FROM todos;

-- name: GetTodo :one
SELECT * FROM todos WHERE id=$1 LIMIT 1;
