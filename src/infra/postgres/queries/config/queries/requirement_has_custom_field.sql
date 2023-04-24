-- requirement_has_custom_field

-- name: SelectCustomFieldIDsAndValuesFromRequirementHasCustomFieldByRequirementID :many
select custom_field_id, value
from requirement_has_custom_field
where requirement_id = @requirement_id;

-- name: SelectValueFromRequirementHasCustomFieldByRequirementIDAndCustomFieldID :one
select value from requirement_has_custom_field
where requirement_id = @requirement_id and custom_field_id = @custom_field_id;

-- name: InsertIntoRequirementHasCustomField :exec
insert into requirement_has_custom_field (requirement_id, custom_field_id, value)
values (@requirement_id, @custom_field_id, @value);

-- name: UpdateValueFromRequirementHasCustomFieldByRequirementIDAndCustomFieldID :exec
update requirement_has_custom_field
set value = @value
where requirement_id = @requirement_id and custom_field_id = @custom_field_id;

-- name: DeleteFromRequirementHasCustomFieldByRequirementIDAndCustomFieldID :exec
delete from requirement_has_custom_field
where requirement_id = @requirement_id and custom_field_id = @custom_field_id;
