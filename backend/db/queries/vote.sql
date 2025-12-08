-- name: GetRecent10VotesByUserId :many
select * from "votes" where "userId" = $1 order by "dateTime" desc limit 10;

-- name: CreateVote :exec
insert into "votes" ("userId", "forCharacterId", "againstCharacterId") values ($1, $2, $3);