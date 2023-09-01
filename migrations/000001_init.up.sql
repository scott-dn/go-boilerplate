set timezone = utc;

create table books (
		id serial primary key,
		name varchar(255) not null,
		author varchar(255) not null,
    description varchar(5000) not null,
		version int not null,
		created_by varchar(255) not null,
		created_at timestamp (6) with time zone default now() not null,
		updated_by varchar(255) not null,
		updated_at timestamp (6) with time zone default now() not null
);
