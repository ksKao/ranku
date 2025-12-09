-- name: CreateLike :exec
insert into "like" ("userId", "characterId") values ($1, $2);

-- name: GetUserLikes :many
select "character".*, "anime"."name" as "anime"
from
    "like"
    join "character" on "character"."id" = "like"."characterId"
    join "anime_character" on "anime_character"."characterId" = "character"."id"
    join "anime" on "anime_character"."animeId" = "anime"."id"
where "like"."userId" = $1;