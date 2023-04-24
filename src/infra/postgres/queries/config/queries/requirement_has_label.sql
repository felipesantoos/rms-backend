-- requirement_has_label

-- name: SelectLabelIDFromRequirementHasLabelByRequirementID :many
select label_id from requirement_has_label where requirement_id = @requirement_id;

-- name: SelectRequirementIDFromRequirementHasLabelByLabelID :many
select requirement_id from requirement_has_label where label_id = @label_id;

-- name: InsertIntoRequirementHasLabel :exec
insert into requirement_has_label (requirement_id, label_id)
values (@requirement_id, @label_id);

-- name: DeleteFromRequirementHasLabelByRequirementIDAndLabelID :exec
delete from requirement_has_label
where requirement_id = @requirement_id and label_id = @label_id;
