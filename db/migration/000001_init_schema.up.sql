create table accounts
(
    id             bigserial
        primary key,
    username       varchar(255)                            not null,
    account_number varchar(255)                            not null
        unique,
    status         boolean                  default true   not null,
    balance        numeric(19, 4)           default 0.0000 not null,
    account_type   varchar(255)                            not null,
    created_at     timestamp with time zone default now()  not null,
    updated_at     timestamp with time zone default now()  not null
);

create table entries
(
    id             bigserial
        primary key,
    account_number varchar(255)                           not null,
    amount         numeric(19, 4)                         not null,
    created_at     timestamp with time zone default now() not null
);

create table transfers
(
    id                  bigserial
        primary key,
    from_account_number varchar(255)                           not null,
    to_account_number   varchar(255)                           not null,
    amount              numeric(19, 4)                         not null,
    created_at          timestamp with time zone default now() not null
);

CREATE INDEX ON "accounts" ("username");

CREATE UNIQUE INDEX ON "accounts" ("account_number");

CREATE INDEX ON "entries" ("account_number");

CREATE INDEX ON "transfers" ("from_account_number");

CREATE INDEX ON "transfers" ("to_account_number");

CREATE INDEX ON "transfers" ("from_account_number", "to_account_number");

COMMENT ON COLUMN "entries"."amount" IS 'can be negative or positive';

COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';
