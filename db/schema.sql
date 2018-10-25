DROP DATABASE "go-exchange-engine";
CREATE DATABASE "go-exchange-engine" ENCODING = "UTF-8";

DROP TABLE IF EXISTS "users";
CREATE TABLE IF NOT EXISTS "users" (
    "id" SERIAL PRIMARY KEY NOT NULL,
    "email" STRING NOT NULL UNIQUE,
    "created_at" TIMESTAMP,
    "updated_at" TIMESTAMP
);

DROP TABLE IF EXISTS "accounts";
CREATE TABLE IF NOT EXISTS "accounts" (
    "id" SERIAL PRIMARY KEY,
    "user_id" INT,
    "currency" STRING,
    "balance" DECIMAL,
    "available_balance" DECIMAL,
    "created_at" TIMESTAMP,
    "updated_at" TIMESTAMP
);

DROP TABLE IF EXISTS "deposits";
CREATE TABLE IF NOT EXISTS "deposits" (
    "id" SERIAL PRIMARY KEY,
    "user_id" INT,
    "currency" STRING,
    "amount" DECIMAL,
    "status" STRING,
    "confirmations" INT,
    "created_at" TIMESTAMP,
    "updated_at" TIMESTAMP
);

DROP TABLE IF EXISTS "withdraws";
CREATE TABLE IF NOT EXISTS "withdraws" (
    "id" SERIAL PRIMARY KEY,
    "user_id" INT,
    "currency" STRING,
    "amount" DECIMAL,
    "status" STRING,
    "created_at" TIMESTAMP,
    "updated_at" TIMESTAMP
);

DROP TABLE IF EXISTS "accounting_entries";
CREATE TABLE IF NOT EXISTS "accounting_entries" (
    "id" SERIAL PRIMARY KEY,
    "entryable_type" STRING,
    "entryable_id" STRING,
    "credit_amount" DECIMAL,
    "credit_account_id" INT,
    "debit_amount" DECIMAL,
    "debit_account_id" INT,
    "currency" STRING,
    "created_at" TIMESTAMP,
    "updated_at" TIMESTAMP
);

DROP TABLE IF EXISTS "orders"
CREATE TABLE IF NOT EXISTS "orders" (
    "id" SERIAL PRIMARY KEY,
    "user_id" INT,
    "side" STRING,
    "type" STRING,
    "pair" STRING,
    "amount" DECIMAL,
    "price" DECIMAL,
    "status" STRING,
    "sequence_no" INT,
    "origin_funds" DECIMAL,
    "left_funds" DECIMAL,
    "done_at" TIMESTAMP,
    "created_at" TIMESTAMP,
    "updated_at" TIMESTAMP
);


DROP TABLE IF EXISTS "trades"
CREATE TABLE IF NOT EXISTS "trades" (
    "id" SERIAL PRIMARY KEY,
    "sequence_no" INT,
    "buyer_id" INT,
    "seller_id" INT,
    "amount" DECIMAL,
    "price" DECIMAL,
    "total_price" DECIMAL,
    "ask_order_id" INT,
    "bid_order_id" INT,
    "pair" STRING,
    "created_at" TIMESTAMP,
    "updated_at" TIMESTAMP
);

