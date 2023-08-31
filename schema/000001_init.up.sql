CREATE TABLE users
(
    id serial not null unique
);

CREATE TABLE segments
(
    id serial not null unique,
    title varchar(255) not null
);

CREATE TABLE users_segments
(
    id serial not null unique,
    user_id int references users (id) on delete cascade not null,
    segment_id int references segments (id) on delete cascade not null,
    created_time timestamp not null default now(),
    ttl timestamp not null default 'infinity'::timestamp without time zone,
    action_type varchar(1) not null
);