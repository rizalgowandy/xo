CREATE TABLE authors (
  author_id INTEGER AUTO_INCREMENT NOT NULL PRIMARY KEY,
  name VARCHAR(255) DEFAULT '' NOT NULL
) ENGINE=InnoDB;

CREATE INDEX authors_name_idx ON authors(name);

CREATE TABLE books (
  book_id INTEGER AUTO_INCREMENT NOT NULL PRIMARY KEY,
  author_id INTEGER NOT NULL REFERENCES authors(author_id),
  isbn VARCHAR(255) DEFAULT '' NOT NULL UNIQUE,
  book_type ENUM('FICTION', 'NONFICTION') DEFAULT 'FICTION' NOT NULL,
  title VARCHAR(255) DEFAULT '' NOT NULL,
  year INTEGER DEFAULT 2000 NOT NULL,
  available DATETIME DEFAULT NOW() NOT NULL,
  tags TEXT DEFAULT '' NOT NULL
) ENGINE=InnoDB;

CREATE INDEX books_title_idx ON books(title, year);

CREATE FUNCTION say_hello(name VARCHAR(255)) RETURNS VARCHAR(255)
  DETERMINISTIC
BEGIN
  RETURN CONCAT('hello ', name)\;
END;
