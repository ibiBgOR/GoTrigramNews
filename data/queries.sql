-- name: drop-database
DROP DATABASE IF EXISTS trigramnews;
-- name: create-database
CREATE DATABASE IF NOT EXISTS trigramnews;
-- name: use-database
USE trigramnews;
-- name: create-titles-table
CREATE TABLE IF NOT EXISTS titles ( 
	id INTEGER,
	title VARCHAR(255), 
	PRIMARY KEY (id)
);
-- name: create-trigrams-table
CREATE TABLE IF NOT EXISTS trigrams ( 
	id INTEGER, 
	trigram CHARACTER(3), 
	UNIQUE (id, trigram), 
	FOREIGN KEY (id) REFERENCES titles(id)
);
-- name: insert-title
INSERT INTO titles (id, title) VALUES(?, ?);
-- name: insert-trigram
INSERT INTO trigrams (id, trigram) VALUES(?, ?);
-- name: select-title
SELECT title FROM titles WHERE id = ? LIMIT 1;
-- name: select-trigram
SELECT id FROM titles WHERE trigram = ?;
