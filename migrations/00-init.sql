create table users (
	id uuid primary key,
	email text not null,
	role int not null,
	passwordHash bytea not null
);
create index idxUsersEmail on users(email);

create table pvzs (
	id uuid primary key,
	registrationDate timestamptz not null,
	city int not null
);

create table receptions (
	id uuid primary key,
	pvzId uuid references pvzs(id) on delete cascade,
	dateTime timestamptz not null,
	status int not null
);
create index idxReceptionsPvzId on receptions(pvzId);

create table products (
	id uuid primary key,
	dateTime timestamptz not null,
	receptionId uuid references receptions(id) on delete cascade,
	type int not null
);
create index idxProductsReceptionId on products(receptionId);
create index idxProductsDateTime on products(dateTime);
