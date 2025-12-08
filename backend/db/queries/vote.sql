-- name: GetRecent10VotesByUserId :many
select * from "votes" where "userId" = $1 order by "dateTime" desc limit 10;