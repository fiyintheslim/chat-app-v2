CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	email TEXT UNIQUE NOT NULL,
	first_name TEXT,
	last_name TEXT,
	username TEXT UNIQUE,
	created_at BIGINT NOT NULL DEFAULT extract(epoch FROM now()),
	updated_at BIGINT NOT NULL DEFAULT extract(epoch FROM now()),
	deleted_at BIGINT DEFAULT NULL,
	bio TEXT,
	password TEXT NOT NULL DEFAULT "",
);