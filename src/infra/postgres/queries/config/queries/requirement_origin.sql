-- requirement_origin

-- name: SelectAllFromRequirementOrigin :many
select * from requirement_origin;

-- name: SelectRequirementOriginByID :one
select * from requirement_origin where id = @id;

-- name: InsertIntoRequirementOrigin :one
insert into requirement_origin (id, name) values (default, @name) returning id;

-- name: UpdateRequirementOriginByID :exec
update requirement_origin set name = @name where id = @id;

-- name: DeleteFromRequirementOriginByID :exec
delete from requirement_origin where id = @id;
