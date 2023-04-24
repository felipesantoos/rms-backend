-- custom_field_type

-- name: SelectAllFromCustomFieldType :many
select * from custom_field_type;

-- name: SelectCustomFieldTypeByID :one
select * from custom_field_type where id = @id;

-- name: InsertIntoCustomFieldType :one
insert into custom_field_type (id, name) values (default, @name)
returning id;

-- name: UpdateCustomFieldTypeByID :exec
update custom_field_type set name = @name where id = @id;

-- name: DeleteFromCustomFieldTypeByID :exec
delete from custom_field_type where id = @id;
