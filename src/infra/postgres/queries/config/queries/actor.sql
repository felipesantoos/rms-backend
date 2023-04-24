-- actor

-- name: SelectAllFromActor :many
select * from actor;

-- name: SelectFromActorByID :one
select * from actor where id = @id;

-- name: InsertIntoActor :one
insert into actor (id, name, description)
values (default, @name, @description)
returning id;

-- name: UpdateActorById :exec
update actor set name = @name, description = @description where id = @id;

-- name: DeleteFromActorById :exec
delete from actor where id = @id;
