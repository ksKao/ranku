-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

create table "like" (
    "userId" text not null references "user" ("id") on delete cascade,
    "characterId" uuid not null references "character" ("id") on delete cascade,
    primary key ("userId", "characterId")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists "like" cascade;
-- +goose StatementEnd