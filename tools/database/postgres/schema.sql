create table project (
    id serial primary key,
    name varchar(200) not null,
    alias varchar(10),
    description text null,
    is_active bool default true,
    created_at timestamp not null default now(),
    updated_at timestamp null,
    deleted_at timestamp null
);

create table requirement_type (
    id serial primary key,
    name varchar(200) not null unique
);

create table requirement_origin (
    id serial primary key,
    name varchar(200) not null unique
);

create table priority (
    id serial primary key,
    level varchar(200) not null unique
);

create table requirement (
    id serial primary key,
    code varchar(20) not null,
    title varchar(200) not null,
    description text null,
    user_story text null,
    type_id int not null references requirement_type (id),
    origin_id int not null references requirement_origin (id),
    priority_id int not null references priority (id),
    project_id int not null references project (id),
    created_at timestamp not null default now(),
    updated_at timestamp null,
    deleted_at timestamp null
);

create table acceptance_criteria (
    id serial primary key,
    requirement_id int not null references requirement (id),
    created_at timestamp not null default now(),
    updated_at timestamp null,
    deleted_at timestamp null
);

create table label (
    id serial primary key,
    name varchar(200) not null unique
);

create table requirement_has_label (
    requirement_id int not null references requirement (id),
    label_id int not null references label (id),
    constraint pk_requirement_has_label primary key (requirement_id, label_id)
);

create table requirement_depends_on_requirement (
    dependent_id int not null references requirement (id),
    prerequisite_id int not null references requirement (id)
);

create table actor (
    id serial primary key,
    name varchar(200) not null,
    description text null
);

create table actor_is_associated_with_requirement (
    actor_id int not null references actor (id),
    requirement_id int not null references requirement (id),
    constraint pk_actor_is_associated_with_requirement primary key (actor_id, requirement_id)
);

create table custom_field_type (
    id serial primary key,
    name varchar(200) not null unique
);

create table custom_field (
    id serial primary key,
    name varchar(200) not null,
    description text null,
    type_id int not null references custom_field_type (id)
);

create table requirement_has_custom_field (
    requirement_id int not null references requirement (id),
    custom_field_id int not null references custom_field (id),
    value text null,
    constraint pk_requirement_has_custom_field primary key (requirement_id, custom_field_id)
);
