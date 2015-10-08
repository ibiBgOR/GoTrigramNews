-- name: drop-database
DROP DATABASE IF EXISTS trigramnews;
-- name: create-database
CREATE DATABASE IF NOT EXISTS trigramnews;
-- name: use-database
USE trigramnews;
-- name: create-titles-table
CREATE TABLE IF NOT EXISTS titles ( 
	id INTEGER NOT NULL AUTO_INCREMENT,
	title VARCHAR(255) NOT NULL, 
	PRIMARY KEY (id)
);
-- name: create-trigrams-table
CREATE TABLE IF NOT EXISTS trigrams ( 
	id INTEGER NOT NULL, 
	trigram CHARACTER(3) NOT NULL, 
	UNIQUE (id, trigram), 
	FOREIGN KEY (id) REFERENCES titles(id)
);
-- name: insert-title
INSERT INTO titles (title) VALUES(?);
-- name: insert-trigram
INSERT INTO trigrams (id, trigram) VALUES(?, ?);
-- name: select-titleids-from-trigram
SELECT id FROM trigrams WHERE trigram = ?;
-- name: select-title
SELECT title FROM titles WHERE id = ? LIMIT 1;
-- name: select-titleid-by-name
SELECT id FROM titles WHERE title = ? LIMIT 1;
