-- name: GetTodos :many
SELECT * FROM todos WHERE todo_list_id = $1 ORDER BY id DESC;

-- name: GetTodo :one
SELECT * FROM todos WHERE id=$1 LIMIT 1;

-- name: UpdateTodoComplete :one
UPDATE todos
SET complete = $2
WHERE id = $1
RETURNING *;

-- name: InsertTodo :one
INSERT INTO todos (title, todo_list_id) VALUES ($1, $2) RETURNING *;

-- name: DeleteTodo :one
DELETE FROM todos
WHERE id=$1
RETURNING id;

-- name: GetTodoLists :many
SELECT id, title, created_at FROM todo_lists ORDER BY title ASC;

