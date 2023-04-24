-- requirement_depends_on_requirement

-- name: SelectDependentIDFromRequirementDependsOnRequirementByPrerequisiteID :many
select dependent_id from requirement_depends_on_requirement where prerequisite_id = @prerequisite_id;

-- name: SelectPrerequisiteIDFromRequirementDependsOnRequirementByDependentID :many
select prerequisite_id from requirement_depends_on_requirement where dependent_id = @dependent_id;

-- name: InsertIntoRequirementDependsOnRequirement :exec
insert into requirement_depends_on_requirement (dependent_id, prerequisite_id) values (@dependent_id, @prerequisite_id);

-- name: DeleteFromRequirementDependsOnRequirementByDependentIDAndPrerequisiteID :exec
delete from requirement_depends_on_requirement where dependent_id = @dependent_id and prerequisite_id = @prerequisite_id;
