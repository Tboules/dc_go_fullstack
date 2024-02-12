-- name: GetSession :one
SELECT * FROM `session`
WHERE token = ? LIMIT 1;

-- name: SaveSession :execresult
INSERT INTO `session` (
  token, user_id, expires_at
) VALUES (
  ?, ?, ?
);

-- name: DeleteSession :exec
DELETE FROM `session`
WHERE token = ?;
