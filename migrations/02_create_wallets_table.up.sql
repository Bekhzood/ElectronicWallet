CREATE TABLE IF NOT EXISTS "wallets" (
  "id"              varchar PRIMARY KEY,
  "number"          int NOT NULL UNIQUE,
  "type"            varchar NOT NULL,
  "balance"         int NOT NULL,
  "user_id"         varchar,
  "status"          varchar DEFAULT 'active',
  "created_at"      timestamp DEFAULT (now()),
  "updated_at"      timestamp DEFAULT (now())
);
