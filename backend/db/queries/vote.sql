-- name: GetUserVoteWithCharacterIds :one
select * from "vote" where "userId" = $1 and (("forCharacterId" = $2 and "againstCharacterId" = $3) or ("forCharacterId" = $3 and "againstCharacterId" = $2)) limit 1;

-- name: CreateVote :exec
insert into "vote" ("userId", "forCharacterId", "againstCharacterId") values ($1, $2, $3);

-- name: GetUserVotes :many
select * from "vote" where "userId" = $1;

-- name: GetTop100VotedCharacters :many
SELECT "character"."id", "character"."name", "character"."image", COALESCE(
        SUM(
            CASE
                WHEN "vote"."forCharacterId" = "character"."id" THEN 1
                ELSE 0
            END
        ), 0
    ) AS for, COALESCE(
        SUM(
            CASE
                WHEN "vote"."againstCharacterId" = "character"."id" THEN 1
                ELSE 0
            END
        ), 0
    ) AS against, COALESCE(
        SUM(
            CASE
                WHEN "vote"."forCharacterId" = "character"."id" THEN 1
                ELSE 0
            END
        ), 0
    ) - COALESCE(
        SUM(
            CASE
                WHEN "vote"."againstCharacterId" = "character"."id" THEN 1
                ELSE 0
            END
        ), 0
    ) AS score
FROM "character"
    LEFT JOIN "vote" ON "vote"."forCharacterId" = "character".id
    OR "vote"."againstCharacterId" = "character"."id"
GROUP BY
    "character"."id"
ORDER BY score DESC
LIMIT 100;