-- priority

-- name: SelectAllFromPriority :many
select * from priority;

-- name: SelectPriorityByID :one
select * from priority where id = @id;

-- name: InsertIntoPriority :one
insert into priority (id, level) values (default, @level) returning id;

-- name: UpdatePriorityByID :exec
update priority set level = @level where id = @id;

-- name: DeleteFromPriorityByID :exec
delete from priority where id = @id;
