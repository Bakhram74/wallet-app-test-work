CREATE TABLE "wallets" (
"wallet_id" uuid DEFAULT gen_random_uuid(),
"balance"  bigint NOT NULL,
"created_at" TIMESTAMPTZ  NOT NULL DEFAULT (now()),
PRIMARY KEY (wallet_id)
);

CREATE INDEX idx_wallet_id ON "wallets" ("wallet_id");
