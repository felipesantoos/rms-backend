-- actor_is_associated_with_requirement

-- name: SelectActorIDFromActorIsAssociatedWithRequirementByRequirementID :many
select actor_id from actor_is_associated_with_requirement where requirement_id = @requirement_id;

-- name: SelectRequirementIDFromActorIsAssociatedWithRequirementByActorID :many
select requirement_id from actor_is_associated_with_requirement where actor_id = @actor_id;

-- name: InsertIntoActorIsAssociatedWithRequirement :exec
insert into actor_is_associated_with_requirement (actor_id, requirement_id) values (@actor_id, @requirement_id);

-- name: DeleteFromActorIsAssociatedWithRequirementByActorIDAndRequirementID :exec
delete from actor_is_associated_with_requirement where actor_id = @actor_id and requirement_id = @requirement_id;
