create database ttms;
use ttms;
CREATE TABLE customer (
                          id INT AUTO_INCREMENT PRIMARY KEY,
                          username VARCHAR(255) NOT NULL,
                          password VARCHAR(255) NOT NULL,
                          email VARCHAR(255) NOT NULL,
                          UNIQUE INDEX unique_username (username),
                          UNIQUE INDEX unique_email (email)
);
CREATE TABLE manager (
                         id INT AUTO_INCREMENT PRIMARY KEY,
                         username VARCHAR(255) NOT NULL,
                         password VARCHAR(255) NOT NULL,
                         email VARCHAR(255) NOT NULL,
                         UNIQUE INDEX unique_username (username),
                         UNIQUE INDEX unique_email (email)
);