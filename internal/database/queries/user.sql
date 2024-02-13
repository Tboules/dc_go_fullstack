-- name: GetUser :one
SELECT * FROM `user`
WHERE id = ? LIMIT 1;

-- name: GetUserByProviderId :one
SELECT * FROM `user`
WHERE `provider_id` = ? LIMIT 1;

-- name: GetUsers :many
SELECT * FROM `user`
ORDER BY `name`;

-- name: CreateNewUser :execresult
INSERT INTO `user` (
  name, email, image, provider_id
) VALUES (
  ?, ?, ?, ?
);

-- name: UpdateUser :execresult
UPDATE `user`
  SET name = ?,
  email = ?,
  provider_id = ?,
  image = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM `user`
WHERE `id` = ?;
