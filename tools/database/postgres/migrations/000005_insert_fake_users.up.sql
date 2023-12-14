COPY account (id, email, first_name, last_name, password, salt, is_active, created_at, updated_at, deleted_at)
    FROM '/fixtures/000005/account.csv'
    DELIMITER ';' CSV HEADER;
