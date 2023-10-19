-- Extensions

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Tables

CREATE TABLE requirement_type (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT requirement_type_id_pk PRIMARY KEY (id)
);

CREATE TABLE requirement_origin (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT requirement_origin_id_pk PRIMARY KEY (id)
);

CREATE TABLE priority (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    level VARCHAR(100) NOT NULL UNIQUE,
    description TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT priority_id_pk PRIMARY KEY (id)
);

CREATE TABLE requirement (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    code SERIAL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    user_story TEXT NOT NULL,
    type_id UUID NOT NULL,
    origin_id UUID NOT NULL,
    priority_id UUID NOT NULL,
    project_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT requirement_pk PRIMARY KEY (id),
    CONSTRAINT requirement_type_id_fk FOREIGN KEY (type_id) REFERENCES requirement_type (id),
    CONSTRAINT requirement_origin_id_fk FOREIGN KEY (origin_id) REFERENCES requirement_origin (id),
    CONSTRAINT requirement_priority_id_fk FOREIGN KEY (priority_id) REFERENCES priority (id),
    CONSTRAINT requirement_project_id_fk FOREIGN KEY (project_id) REFERENCES project (id)
);

-- Datas

COPY requirement_type (id, name, description, created_at, updated_at)
    FROM '/fixtures/000002/requirement_type.csv'
    DELIMITER ';' CSV HEADER;

COPY requirement_origin (id, name, description, created_at, updated_at)
    FROM '/fixtures/000002/requirement_origin.csv'
    DELIMITER ';' CSV HEADER;

COPY priority (id, level, description, created_at, updated_at)
    FROM '/fixtures/000002/priority.csv'
    DELIMITER ';' CSV HEADER;

COPY requirement (id, title, description, user_story, type_id, origin_id, priority_id, project_id, created_at, updated_at)
    FROM '/fixtures/000002/requirement.csv'
    DELIMITER ';' CSV HEADER;
