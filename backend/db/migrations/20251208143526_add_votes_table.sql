-- +goose Up
-- +goose StatementBegin
create table "votes" (
    "userId" text not null references "user" ("id") on delete cascade,
    "forCharacterId" uuid not null references "character" ("id") on delete cascade,
    "againstCharacterId" uuid not null references "character" ("id") on delete cascade,
    "dateTime" timestamptz not null default now (),
    primary key (
        "userId",
        "forCharacterId",
        "againstCharacterId"
    )
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table "votes"
-- +goose StatementEnd