CREATE TABLE IF NOT EXISTS waste_types (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "name" VARCHAR(25) UNIQUE,
    "point" DECIMAL(7, 2),
    "description" TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL
);