-- requirement

-- name: SelectAllFromRequirement :many
select * from requirement;

-- name: SelectRequirementByID :one
select * from requirement where id = @id;

-- name: InsertIntoRequirement :one
insert into requirement (id, code, title, description, user_story, type_id, origin_id, priority_id, project_id, created_at, updated_at, deleted_at)
values (default, @code, @title, @description, @user_story, @type_id, @origin_id, @priority_id, @project_id, now(), null, null)
returning id;

-- name: UpdateRequirementByID :exec
update requirement
set
    code = @code,
    title = @title,
    description = @description,
    user_story = @user_story,
    type_id = @type_id,
    origin_id = @origin_id,
    priority_id = @priority_id,
    project_id = @project_id,
    updated_at = now()
where id = @id;

-- name: DeleteFromRequirementByID :exec
update requirement
set deleted_at = now()
where id = @id;
