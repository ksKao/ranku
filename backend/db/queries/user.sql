-- name: GetUserById :one
select * from "user" where "id" = $1 limit 1;