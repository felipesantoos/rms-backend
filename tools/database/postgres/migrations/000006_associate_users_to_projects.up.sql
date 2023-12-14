COPY project_contains_user (user_id, project_id, created_at, updated_at, deleted_at)
    FROM '/fixtures/000006/project_contains_user.csv'
    DELIMITER ';' CSV HEADER;
