-- name: ListTodos :many
SELECT id, title, completed, created_at, updated_at
FROM todos
ORDER BY created_at DESC;

-- name: CreateTodo :exec
INSERT INTO todos (id, title, completed, created_at, updated_at) VALUES ($1, $2, $3, $4, $5);

-- name: UpdateTodoCompleted :exec
UPDATE todos
SET completed = $2, updated_at = NOW()
WHERE id = $1;

-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = $1;