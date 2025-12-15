-- name: GetCharacterByAnilistId :one
select * from "character" where "anilistId" = $1 limit 1;

-- name: CreateCharacter :one
insert into
    "character" (
        "image",
        "name",
        "anilistId",
        "birthYear",
        "birthMonth",
        "birthDay",
        "bloodType",
        "age",
        "description",
        "gender"
    )
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) returning *;

-- name: UpdateCharacterById :exec
update "character" set "image" = $1, "name" = $2, "birthYear" = $3, "birthMonth" = $4, "birthDay" = $5, "bloodType" = $6, "age" = $7, "description" = $8, "gender" = $9 where "id" = $10;

-- name: GetAnimeCharacterRelationByIds :one
select * from "anime_character" where "animeId" = $1 and "characterId" = $2 limit 1;

-- name: LinkCharacterToAnime :exec
insert into "anime_character" ("animeId", "characterId") values ($1, $2);

-- name: SearchCharacter :many
select distinct on ("character"."id") "character".*, "anime"."name" as "anime"
from
    "character"
    join "anime_character" on "anime_character"."characterId" = "character"."id"
    join "anime" on "anime"."id" = "anime_character"."animeId"
where "anime"."name" ilike $1 or "character"."name" ilike $1
order by "character"."id", "character"."anilistId";

-- name: GetCharacterById :many
select
    "character".*,
    "anime"."name" as "anime",
    count("like"."userId") as "likes",
    case
        when exists (
            select 1
            from "like"
            where
                "like"."characterId" = "character"."id"
                and "like"."userId" = sqlc.narg('userId')::text
        ) then true
        else false
    end as "liked"
from
    "character"
    join "anime_character" on "anime_character"."characterId" = "character"."id"
    join "anime" on "anime"."id" = "anime_character"."animeId"
    left join "like" on "like"."characterId" = "character"."id"
where
    "character"."id" = $1
group by
    "character"."id",
    "anime"."id";

-- name: GetAllCharactersByRandomOrder :many
select * from "character" order by random();