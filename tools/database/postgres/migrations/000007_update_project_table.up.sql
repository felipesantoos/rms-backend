ALTER TABLE project ADD COLUMN created_by UUID;

-- Datas

TRUNCATE project RESTART IDENTITY CASCADE;

COPY project (id, name, alias, description, is_active, created_by, created_at, updated_at, deleted_at)
    FROM '/fixtures/000007/project.csv'
    DELIMITER ';' CSV HEADER;

COPY requirement (id, title, description, user_story, type_id, origin_id, priority_id, project_id, created_at, updated_at)
    FROM '/fixtures/000002/requirement.csv'
    DELIMITER ';' CSV HEADER;

COPY project_contains_user (user_id, project_id, created_at, updated_at, deleted_at)
    FROM '/fixtures/000006/project_contains_user.csv'
    DELIMITER ';' CSV HEADER;
