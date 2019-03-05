CREATE TABLE owners (
	id serial PRIMARY KEY,
	name varchar(255) NOT NULL
);

CREATE TABLE books (
	id serial PRIMARY KEY,
	name varchar(300) NOT NULL,
	owners_id integer NOT NULL,
	CONSTRAINT books_owners_id_fkey FOREIGN KEY (owners_id) REFERENCES owners (id) ON UPDATE NO ACTION ON DELETE NO ACTION
);