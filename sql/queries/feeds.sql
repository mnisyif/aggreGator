-- name: CreateFeed :one
INSERT INTO feeds(id, created_at, updated_at, name, url, user_id)
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeedByUserID :many
SELECT * FROM feeds WHERE user_id = $1;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedByURL :one
SELECT * FROM feeds WHERE url = $1;

-- name: MarkFeedFetched :one
UPDATE feeds
SET updated_at = NOW(), last_fetched_at = NOW()
WHERE id = $1
RETURNING *;

-- name: ResetFeeds :exec
DELETE FROM feeds;
