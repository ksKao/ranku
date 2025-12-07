-- name: GetAnimeByAnidbId :one
select * from "anime" where "anilistId" = $1 limit 1;

-- name: CreateAnime :one
insert into "anime" ("name", "anilistId") values ($1, $2) returning *;

-- name: UpdateAnimeNameById :one
update "anime" set "name" = $1 where "anime"."id" = $2 returning *;