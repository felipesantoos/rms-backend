-- Extensions

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Tables

CREATE TABLE IF NOT EXISTS account (
    id UUID NOT NULL
    CONSTRAINT df_account_id DEFAULT uuid_generate_v4()
    CONSTRAINT pk_account_id PRIMARY KEY,
    email VARCHAR(256) NOT NULL
    CONSTRAINT uk_account_email UNIQUE,
    first_name VARCHAR(200) NOT NULL,
    last_name VARCHAR(200) NOT NULL,
    password TEXT NOT NULL,
    salt VARCHAR(1024) NOT NULL,
    is_active BOOL NOT NULL
    CONSTRAINT df_account_is_active DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL
    CONSTRAINT df_account_created_at DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL
    CONSTRAINT df_account_updated_at DEFAULT NOW(),
    deleted_at TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS project_contains_user (
    user_id UUID NOT NULL,
    project_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    CONSTRAINT project_contains_user_pk PRIMARY KEY (user_id, project_id),
    CONSTRAINT user_id_fk FOREIGN KEY (user_id) REFERENCES account (id),
    CONSTRAINT project_id_fk FOREIGN KEY (project_id) REFERENCES project (id)
);
