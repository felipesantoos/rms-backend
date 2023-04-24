-- label

-- name: SelectAllFromLabel :many
select * from label;

-- name: SelectLabelByID :one
select * from label where id = @id;

-- name: InsertIntoLabel :one
insert into label (id, name) values (default, @name) returning id;

-- name: UpdateLabelByID :exec
update label set name = @name where id = @id;

-- name: DeleteFromLabelByID :exec
delete from label where id = @id;
