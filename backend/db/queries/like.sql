-- name: CreateLike :exec
insert into "like" ("userId", "characterId") values ($1, $2);

-- name: GetUserLikes :many
select
  "character".*
from
  "like"
  join "character" on "character"."id" = "like"."characterId"
where
  "like"."userId" = $1;

-- name: CheckLikeExists :one
select exists (
    select 1 
    from "like"
    where "userId" = $1 and "characterId" = $2
);

-- name: DeleteLike :exec
delete from "like" where "userId" = $1 and "characterId" = $2;