
create database synapse;

create table crawlers (
    id serial not null,
    created_at timestamp not null default now(),
    reference varchar(36) not null,
    name varchar(100) not null,
    certificate varchar(500) not null,
    primary key (id)
);
