-- project

-- name: SelectAllFromProject :many
select * from project;

-- name: SelectProjectByID :one
select * from project where id = @id;

-- name: InsertIntoProject :one
insert into project (id, name, alias, description, is_active, created_at, updated_at, deleted_at)
values (default, @name, @alias, @description, @is_active, now(), null, null)
returning id;

-- name: UpdateProjectByID :exec
update project
set
    name = @name,
    alias = @alias,
    description = @description,
    is_active = @is_active,
    updated_at = now()
where id = @id;

-- name: DeleteFromProjectByID :exec
update project
set deleted_at = now()
where id = @id;
