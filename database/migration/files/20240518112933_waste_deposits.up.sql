CREATE TABLE IF NOT EXISTS waste_deposits (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    waste_type_id UUID NOT NULL,
    quantity INTEGER NOT NULL,
    "description" TEXT,
    "status" VARCHAR(11) CHECK ("status" IN ('confirmed', 'unconfirmed')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    created_by UUID NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (created_by) REFERENCES users (id),
    FOREIGN KEY (waste_type_id) REFERENCES waste_types (id)
);