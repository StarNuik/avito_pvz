create table users (
	id uuid primary key default gen_random_uuid(),
	email text not null,
	role text not null,
	passwordHash bytea not null
);

create table pvzs (
	id uuid primary key default gen_random_uuid(),
	registrationDate timestamp not null default now(),
	city text not null
);

create table receptions (
	id uuid primary key default gen_random_uuid(),
	pvzId uuid references pvzs(id) on delete cascade,
	dateTime timestamp not null default now(),
	status text not null
);

create table products (
	id uuid primary key default gen_random_uuid(),
	dateTime timestamp not null default now(),
	receptionId uuid references receptions(id) on delete cascade,
	type text not null
);
