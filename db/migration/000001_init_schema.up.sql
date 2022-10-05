CREATE TABLE "accounts"
(
    "id"           bigserial PRIMARY KEY,
    "user_id"      bigint         NOT NULL,
    "status"       boolean        NOT NULL default true,
    "balance"      numeric(19, 4) NOT NULL default 0.0000,
    "account_type" varchar(255)   NOT NULL,
    "created_at"   timestamptz    NOT NULL DEFAULT (now()),
    "updated_at"   timestamptz    NOT NULL DEFAULT (now())
);

CREATE TABLE "entries"
(
    "id"         bigserial PRIMARY KEY,
    "account_id" bigint         NOT NULL,
    "amount"     numeric(19, 4) NOT NULL,
    "created_at" timestamptz    NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers"
(
    "id"              bigserial PRIMARY KEY,
    "from_account_id" bigint         NOT NULL,
    "to_account_id"   bigint         NOT NULL,
    "amount"          numeric(19, 4) NOT NULL,
    "created_at"      timestamptz    NOT NULL DEFAULT (now())
);

ALTER TABLE "entries"
    ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers"
    ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers"
    ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");

CREATE INDEX ON "accounts" ("user_id");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "transfers" ("from_account_id");

CREATE INDEX ON "transfers" ("to_account_id");

CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "entries"."amount" IS 'can be negative or positive';

COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';
