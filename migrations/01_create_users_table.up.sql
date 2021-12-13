CREATE TABLE IF NOT EXISTS "users" (
  "id"              varchar PRIMARY KEY,
  "username"        varchar NOT NULL,
  "password"        varchar NOT NULL,
  "created_at"      timestamp DEFAULT (now()),
  "updated_at"      timestamp DEFAULT (now())
);
