package db

const (
	CreateUsersTable = `create table if not exists users (
	id integer primary key autoincrement unique,
	name text not null,
	surname text not null,
	phone integer not null,
	email text not null,
	role text not null,
	login text not null unique,
	password text not null,
	remove boolean default false)`
	CreateCatalogsTable = `create table if not exists catalog (
	id integer primary key autoincrement unique,
	name text not null unique,
	remove boolean not null default false)`
	CreateProductsTable = `create table if not exists products (
	id integer primary key autoincrement unique,
	category text references catalog(name) not null,
	name text not null unique,
	photo text, 
	cost integer not null,
	status text not null)`
	CreateOrdersTable = `create table if not exists orders (
	id integer primary key autoincrement,
	user_id integer not null,
	ordered_date text not null,
	updated_date text,
	product_name text not null,
	address text not null,
	status text not null)`
)
