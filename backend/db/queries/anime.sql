-- name: GetAnimeByAnidbId :one
select * from "anime" where "anilistId" = $1 limit 1;

-- name: CreateAnime :one
insert into "anime" ("name", "anilistId") values ($1, $2) returning *;

-- name: UpdateAnimeNameById :one
update "anime" set "name" = $1 where "anime"."id" = $2 returning *;

-- name: GetAnimeNameByCharacterId :one
select "anime"."name"
from "anime"
  join "anime_character" on "anime_character"."animeId" = "anime"."id"
  join "character" on "anime_character"."characterId" = "character"."id"
where "character"."id" = $1
order by "anime"."anilistId"
limit 1;