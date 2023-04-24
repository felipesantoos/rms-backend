-- custom_field

-- name: SelectAllFromCustomField :many
select * from custom_field;

-- name: SelectCustomFieldByID :one
select * from custom_field where id = @id;

-- name: InsertIntoCustomField :one
insert into custom_field (id, name, description, type_id) values (default, @name, @description, @type_id)
returning id;

-- name: UpdateCustomFieldByID :exec
update custom_field set name = @name, description = @description, type_id = @type_id where id = @id;

-- name: DeleteFromCustomFieldByID :exec
delete from custom_field where id = @id;
