CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE project (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    alias VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    is_active BOOL NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    CONSTRAINT project_id_pk PRIMARY KEY (id)
);
