CREATE TABLE IF NOT EXISTS wallets (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    "name" VARCHAR(50) NOT NULL UNIQUE,
    amount DECIMAL(10, 2) DEFAULT 0.00,
    "description" TEXT,
    "type" VARCHAR(8) CHECK ("type" IN ('master', 'extension')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);