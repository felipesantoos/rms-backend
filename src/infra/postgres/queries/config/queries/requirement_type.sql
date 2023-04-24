-- requirement_type

-- name: SelectAllFromRequirementType :many
select * from requirement_type;

-- name: SelectFromRequirementTypeByID :one
select * from requirement_type where id = @id;

-- name: InsertIntoRequirementType :one
insert into requirement_type (id, name) values (default, @name) returning id;

-- name: UpdateRequirementTypeByID :exec
update requirement_type
set name = @name
where id = @id;

-- name: DeleteFromRequirementTypeByID :exec
delete from requirement_type where id = @id;
