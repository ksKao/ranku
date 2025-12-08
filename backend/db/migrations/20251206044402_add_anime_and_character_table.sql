-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table "anime" (
    "id" uuid not null default uuid_generate_v4 () primary key,
    "name" text not null,
    "anilistId" int not null unique
);

create table "character" (
    "id" uuid not null default uuid_generate_v4 () primary key,
    "image" text not null,
    "name" text not null,
    "anilistId" int not null,
    "birthYear" int,
    "birthMonth" int,
    "birthDay" int,
    "bloodType" text,
    "age" text,
    "description" text,
    "gender" text
);

create table "anime_character" (
    "animeId" uuid not null references "anime" ("id") on delete cascade,
    "characterId" uuid not null references "character" ("id") on delete cascade,
    primary key ("animeId", "characterId")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists "anime_character" cascade;

-- Drop the character table
drop table if exists "character" cascade;

-- Drop the anime table
drop table if exists "anime" cascade;

-- Remove the extension if it was added in this migration
drop extension if exists "uuid-ossp";
-- +goose StatementEnd