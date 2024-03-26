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
use ttms;
CREATE TABLE Movies (
                        movie_id INT AUTO_INCREMENT PRIMARY KEY,
                        title VARCHAR(255),
                        director VARCHAR(255),
                        actors TEXT,
                        genre VARCHAR(255),
                        duration VARCHAR(255),
                        language VARCHAR(255),
                        release_date VARCHAR(255),
                        rating Float,
                        description TEXT
);
ALTER TABLE Movies
    ADD CONSTRAINT UNIQUE (title);
CREATE TABLE Screens (
                         screen_id INT AUTO_INCREMENT PRIMARY KEY,
                         screen_name VARCHAR(255),
                         screen_capacity INT
);
ALTER TABLE Screens
    ADD CONSTRAINT UNIQUE (screen_name);
CREATE TABLE Showtimes (
                           showtime_id INT AUTO_INCREMENT PRIMARY KEY,
                           movie_name varchar(255),
                           screen_name varchar(255),
                           showtime_start TIMESTAMP,
                           showtime_end TIMESTAMP,
                           FOREIGN KEY (movie_name) REFERENCES Movies(title),
                           FOREIGN KEY (screen_name) REFERENCES Screens(screen_name)
);
CREATE TABLE Tickets (
                         ticket_id INT AUTO_INCREMENT PRIMARY KEY,
                         showtime_name varchar(255),
                         seat_number INT,
                         price float,
                         customer_name varchar(255)
);
CREATE TABLE Ratings (
                         user_name varchar(255),
                         movie_name varchar(255),
                         score FLOAT,
                         PRIMARY KEY (user_name, movie_name),
                         FOREIGN KEY (user_name) REFERENCES customer(username),
                         FOREIGN KEY (movie_name) REFERENCES Movies(title)
);
CREATE TABLE Comments (
                          comment_id INT AUTO_INCREMENT PRIMARY KEY,
                          user_name varchar(255),
                          movie_name varchar(255),
                          comment TEXT NOT NULL,
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          FOREIGN KEY (user_name) REFERENCES customer(username),
                          FOREIGN KEY (movie_name) REFERENCES Movies(title)
)
ALTER TABLE Movies ADD url varchar(255);