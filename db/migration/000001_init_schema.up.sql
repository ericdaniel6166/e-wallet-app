CREATE TABLE "accounts"
(
    "id"             bigserial PRIMARY KEY,
    "username"       varchar(255)   NOT NULL,
    "account_number" varchar(255)   NOT NULL,
    "status"         boolean        NOT NULL default true,
    "balance"        numeric(19, 4) NOT NULL default 0.0000,
    "account_type"   varchar(255)   NOT NULL,
    "created_at"     timestamptz    NOT NULL DEFAULT (now()),
    "updated_at"     timestamptz    NOT NULL DEFAULT (now())
);

CREATE TABLE "entries"
(
    "id"             bigserial PRIMARY KEY,
    "account_number" varchar(255)   NOT NULL,
    "amount"         numeric(19, 4) NOT NULL,
    "created_at"     timestamptz    NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers"
(
    "id"                  bigserial PRIMARY KEY,
    "from_account_number" varchar(255)   NOT NULL,
    "to_account_number"   varchar(255)   NOT NULL,
    "amount"              numeric(19, 4) NOT NULL,
    "created_at"          timestamptz    NOT NULL DEFAULT (now())
);

CREATE INDEX ON "accounts" ("username");

CREATE UNIQUE INDEX ON "accounts" ("account_number");

CREATE INDEX ON "entries" ("account_number");

CREATE INDEX ON "transfers" ("from_account_number");

CREATE INDEX ON "transfers" ("to_account_number");

CREATE INDEX ON "transfers" ("from_account_number", "to_account_number");

COMMENT ON COLUMN "entries"."amount" IS 'can be negative or positive';

COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';
