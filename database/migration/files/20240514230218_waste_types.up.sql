CREATE TABLE IF NOT EXISTS waste_types (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "name" VARCHAR(25) NOT NULL,
    "point" DECIMAL(7, 2) NOT NULL,
    "description" TEXT,
    "version_id" UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    created_by UUID NOT NULL,
    updated_by UUID DEFAULT NULL,
    deleted_by UUID DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS wt_versions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    product_id UUID,
    "name" VARCHAR(25) NOT NULL,
    "point" DECIMAL(7, 2) NOT NULL,
    "description" TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by UUID NOT NULL
)