create table sessions
(
    id                 bigserial
        primary key,
    username           varchar(255)                           not null,
    refresh_token_uuid uuid                                   not null
        unique,
    refresh_token      varchar                                not null
        unique,
    user_agent         varchar(255)                           not null,
    client_ip          varchar(255)                           not null,
    status             boolean                  default true  not null,
    expires_at         timestamp with time zone default now() not null,
    created_at         timestamp with time zone default now() not null,
    updated_at         timestamp with time zone default now() not null
);

create unique index on sessions (refresh_token_uuid);
create index on sessions (username,client_ip,user_agent);