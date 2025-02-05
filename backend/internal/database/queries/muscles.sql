-- name: GetMuscle :one
SELECT muscle_name, muscle_group
FROM muscles
WHERE muscle_name = $1;

-- name: CreateMuscle :one
INSERT INTO muscles (muscle_name, muscle_group)
VALUES ($1, $2)
RETURNING *;

-- name: DeleteMuscle :one
DELETE FROM muscles
WHERE muscle_name = $1
RETURNING *;

-- name: DeleteAllMuscles :exec
DELETE FROM muscles;