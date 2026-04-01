-- name: GetTodos :many
SELECT * FROM todos;

-- name: GetTodo :one
SELECT * FROM todos WHERE id=$1 LIMIT 1;

-- name: UpdateTodoComplete :one
UPDATE todos
SET complete = $2
WHERE id = $1
RETURNING *;
