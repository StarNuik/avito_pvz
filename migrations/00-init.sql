create table users (
	id uuid primary key default gen_random_uuid(),
	email text not null,
	role int not null,
	passwordHash bytea not null
);
create index idxUsersEmail on users(email);

create table pvzs (
	id uuid primary key default gen_random_uuid(),
	registrationDate timestamp not null default now(),
	city int not null
);
--- TODO idx date?

create table receptions (
	id uuid primary key default gen_random_uuid(),
	pvzId uuid references pvzs(id) on delete cascade,
	dateTime timestamp not null default now(),
	status int not null
);
--- todo idx date?
--- todo idx pvzId?

create table products (
	id uuid primary key default gen_random_uuid(),
	dateTime timestamp not null default now(),
	receptionId uuid references receptions(id) on delete cascade,
	type int not null
);
--- todo idx receptionId?
--- todo idx date?
