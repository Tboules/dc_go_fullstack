-- name: GetUser :one
SELECT * FROM `user`
WHERE id = ? LIMIT 1;

-- name: GetUsers :many
SELECT * FROM `user`
ORDER BY `name`;

-- name: CreateNewUser :execresult
INSERT INTO `user` (
  name, email, image
) VALUES (
  ?, ?, ?
);

-- name: UpdateUser :execresult
UPDATE `user`
  SET name = ?,
  email = ?,
  email_verified = ?,
  image = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM `user`
WHERE `id` = ?;
