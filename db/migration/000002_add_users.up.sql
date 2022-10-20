create table users
(
    id              bigserial
        primary key,
    username        varchar(255)                           not null
        unique,
    status          boolean                  default true  not null,
    role            int                                    not null,
    hashed_password varchar(255)                           not null,
    full_name       varchar(255)                           not null,
    email           varchar(255)                           not null
        unique,
    updated_at      timestamp with time zone default now() not null,
    created_at      timestamp with time zone default now() not null
);

CREATE UNIQUE INDEX ON "users" ("username");

CREATE UNIQUE INDEX ON "users" ("email");