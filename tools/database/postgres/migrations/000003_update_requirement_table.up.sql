ALTER TABLE requirement ADD COLUMN deleted_at TIMESTAMP DEFAULT NULL;
ALTER TABLE requirement ADD CONSTRAINT requirement_code_and_project_id_uk UNIQUE (code, project_id);
