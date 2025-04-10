create table users (
	id uuid primary key,
	email text not null,
	role int not null,
	passwordHash bytea not null
);
create index idxUsersEmail on users(email);

create table pvzs (
	id uuid primary key,
	registrationDate timestamp with time zone not null,
	city int not null
);
--- TODO idx date?

create table receptions (
	id uuid primary key,
	pvzId uuid references pvzs(id) on delete cascade,
	dateTime timestamp with time zone not null,
	status int not null
);
--- todo idx date?
--- todo idx pvzId?

create table products (
	id uuid primary key,
	dateTime timestamp with time zone not null,
	receptionId uuid references receptions(id) on delete cascade,
	type int not null
);
create index idxProductsReceptionId on products(receptionId);
--- todo idx receptionId?
--- todo idx date?
