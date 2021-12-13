CREATE TABLE IF NOT EXISTS "transactions_history" (
  "id"                     varchar PRIMARY KEY,
  "sum"                    int NOT NULL,
  "sender_wallet_number"   int NOT NULL,
  "receiver_wallet_number" int NOT NULL,
  "created_at"             timestamp DEFAULT (now())
);
