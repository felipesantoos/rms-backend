-- Extensions

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Tables

CREATE TABLE project (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    alias VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    is_active BOOL NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    CONSTRAINT project_id_pk PRIMARY KEY (id)
);

-- Datas

COPY project (id, name, alias, description, is_active, created_at, updated_at, deleted_at)
    FROM '/fixtures/000001/project.csv'
    DELIMITER ';' CSV HEADER;
