-- name: GetRooms :many
SELECT
    "id",
    "theme"
FROM
    rooms;

-- name: GetRoomByID :one
SELECT
    "id",
    "theme"
FROM
    rooms
WHERE
    id = $1;

-- name: CreateRoom :one
INSERT INTO rooms ("theme")
    VALUES ($1)
RETURNING
    "id", "theme";

-- name: DeleteRoom :exec
DELETE FROM rooms
WHERE id = $1;

-- name: UpdateRoom :one
UPDATE
    rooms
SET
    "theme" = $2
WHERE
    id = $1
RETURNING
    "id",
    "theme";

