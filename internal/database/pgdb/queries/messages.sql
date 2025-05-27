-- name: GetAllMessages :many
SELECT
    "id",
    "room_id",
    "message",
    "reaction_count",
    "answered"
FROM
    messages;

-- name: GetMessageByID :one
SELECT
    "id",
    "room_id",
    "message",
    "reaction_count",
    "answered"
FROM
    messages
WHERE
    id = $1;

-- name: GetMessagesByRoomID :many
SELECT
    "id",
    "room_id",
    "message",
    "reaction_count",
    "answered"
FROM
    messages
WHERE
    room_id = $1;

-- name: CreateMessage :one
INSERT INTO messages ("room_id", "message")
    VALUES ($1, $2)
RETURNING
    "id", "room_id", "message", "reaction_count", "answered";

-- name: DeleteMessage :exec
DELETE FROM messages
WHERE id = $1;

-- name: UpdateMessage :one
UPDATE
    messages
SET
    "message" = $2,
    "reaction_count" = $3,
    "answered" = $4
WHERE
    id = $1
RETURNING
    "id",
    "room_id",
    "message",
    "reaction_count",
    "answered";

-- name: ReactToMessage :one
UPDATE
    messages
SET
    "reaction_count" = "reaction_count" + 1
WHERE
    id = $1
RETURNING
    "id",
    "room_id",
    "message",
    "reaction_count",
    "answered";

-- name: RemoveReactionFromMessage :one
UPDATE
    messages
SET
    "reaction_count" = "reaction_count" - 1
WHERE
    id = $1
RETURNING
    "id",
    "room_id",
    "message",
    "reaction_count",
    "answered";

-- name: MarkMessageAsAnswered :one
UPDATE
    messages
SET
    "answered" = TRUE
WHERE
    id = $1
RETURNING
    "id",
    "room_id",
    "message",
    "reaction_count",
    "answered";

