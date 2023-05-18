-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2023-05-18T16:15:09.392Z

CREATE TABLE "users" (
  "username" bigserial PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "profile_image" varchar NOT NULL,
  "phone_number" varchar NOT NULL,
  "push_notification_token" varchar NOT NULL,
  "is_used" bool NOT NULL,
  "secret_code" varchar NOT NULL,
  "otp_created_at" timestamptz NOT NULL DEFAULT (now()),
  "expired_at" timestamptz NOT NULL DEFAULT (now()),
  "password_created_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "account_number" bigint UNIQUE NOT NULL,
  "currency" varchar NOT NULL,
  "bank_name" varchar NOT NULL,
  "swift_code" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "cards" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "card_number" bigint UNIQUE NOT NULL,
  "currency" varchar NOT NULL,
  "cvv" int UNIQUE NOT NULL,
  "expired_at" timestamptz NOT NULL,
  "billing_address" varchar NOT NULL,
  "zip_code" int UNIQUE NOT NULL,
  "card_scheme" varchar NOT NULL
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint,
  "account_number" bigint,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transactions" (
  "id" bigserial,
  "transaction_id" bigserial,
  "from_account_number" bigint,
  "to_account_number" bigint,
  "transaction_type" varchar NOT NULL,
  "status" varchar NOT NULL,
  "exchanage_rate" varchar NOT NULL,
  "currency_pair" varchar NOT NULL,
  "transaction_status" bool NOT NULL,
  "amount" bigint NOT NULL,
  "hashed_transaction_pin" varchar NOT NULL,
  "transaction_pin_created_at" varchar NOT NULL,
  "transaction_fees" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("id", "transaction_id")
);

CREATE INDEX ON "accounts" ("owner");

CREATE UNIQUE INDEX ON "accounts" ("owner", "currency");

CREATE INDEX ON "cards" ("owner");

CREATE UNIQUE INDEX ON "cards" ("owner", "currency");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "entries" ("account_number");

CREATE INDEX ON "transactions" ("from_account_number");

CREATE INDEX ON "transactions" ("to_account_number");

CREATE INDEX ON "transactions" ("from_account_number", "to_account_number");

COMMENT ON COLUMN "accounts"."currency" IS 'must be positive';

COMMENT ON COLUMN "entries"."amount" IS 'could be positive or negative';

COMMENT ON COLUMN "transactions"."amount" IS 'must be positive';

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

ALTER TABLE "cards" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

ALTER TABLE "cards" ADD FOREIGN KEY ("currency") REFERENCES "accounts" ("currency");

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "entries" ADD FOREIGN KEY ("account_number") REFERENCES "accounts" ("account_number");

ALTER TABLE "transactions" ADD FOREIGN KEY ("from_account_number") REFERENCES "accounts" ("account_number");

ALTER TABLE "transactions" ADD FOREIGN KEY ("to_account_number") REFERENCES "accounts" ("account_number");
