-- name: CreateFeedFollow :many 
INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
