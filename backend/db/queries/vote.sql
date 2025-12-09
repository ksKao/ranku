-- name: GetUserVoteWithCharacterIds :one
select * from "vote" where "userId" = $1 and (("forCharacterId" = $2 and "againstCharacterId" = $3) or ("forChracterId" = $3 and "andCharacterId" = $2)) limit 1;

-- name: CreateVote :exec
insert into "vote" ("userId", "forCharacterId", "againstCharacterId") values ($1, $2, $3);

-- name: GetUserVotes :many
select * from "vote" where "userId" = $1;