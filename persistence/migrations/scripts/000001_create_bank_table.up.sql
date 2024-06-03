create extension if not exists "pgcrypto";

create table if not exists banks (
    id serial primary key,
    code varchar(50) not null,
    url varchar(50) not null,

    created_at timestamp without time zone default (now() at time zone 'utc'),
    updated_at timestamp without time zone default (now() at time zone 'utc'),
    deleted_at timestamp without time zone default (now() at time zone 'utc'),

    unique(url),
    CHECK(url <> ''),
    unique(code),
    CHECK(code <> '')
);

