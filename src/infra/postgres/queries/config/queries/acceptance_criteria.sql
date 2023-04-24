-- acceptance_criteria

-- name: SelectAllFromAcceptanceCriteria :many
select * from acceptance_criteria;

-- name: SelectAcceptanceCriteriaByID :one
select * from acceptance_criteria where id = @id;

-- name: InsertIntoAcceptanceCriteria :one
insert into acceptance_criteria (id, requirement_id, created_at, updated_at, deleted_at)
values (default, @requirement_id, now(), null, null)
returning id;

-- name: UpdateAcceptanceCriteriaByID :exec
update acceptance_criteria
set
    requirement_id = @requirement_id,
    updated_at = now()
where id = @id;

-- name: DeleteAcceptanceCriteriaByID :exec
update acceptance_criteria
set deleted_at = now()
where id = @id;
