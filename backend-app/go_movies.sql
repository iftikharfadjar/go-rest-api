CREATE DATABASE go_movies_db

CREATE TABLE genres (
id INT NOT NULL AUTO_INCREMENT,
genre_name VARCHAR(20),
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
CONSTRAINT genres_pkey PRIMARY KEY (id)
);


CREATE TABLE movies (
    id INT NOT NULL AUTO_INCREMENT,
    title VARCHAR(200),
    description TEXT,
    year INT,
    release_date DATE,
    runtime INT,
    rating INT,
    mpaa_rating VARCHAR(200),
    created_at TIMESTAMP  DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT movies_pkey PRIMARY KEY (id)
);


CREATE TABLE movies_genres (
    id INT NOT NULL AUTO_INCREMENT,
    movie_id INT,
    genre_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT movies_genres_pkey PRIMARY KEY (id),
	CONSTRAINT fk_movie_genries_genre_id FOREIGN KEY (genre_id) REFERENCES genres(id),
	CONSTRAINT fk_movie_genries_movie_id FOREIGN KEY (movie_id) REFERENCES movies(id)
);



INSERT INTO genres (genre_name) VALUES 
("Drama"),
("Crime"),
("Action"),
("Comic Book"),
("Sci-Fi"),
("Mystery"),
("Adventure"),
("Comedy"),
("Romance");
	
INSERT INTO movies  (title, description, year, release_date, runtime, rating, mpaa_rating) VALUES
("The Shawshank Redemption",	"Two imprisoned men bond over a number of years",	1994,	"1994-10-14"	,142,	5,	"R"),
("The Godfather",	"The aging patriarch of an organized crime dynasty transfers control to his son",	1972,	"1972-03-24"	,175,	5,	"R"),				
("American Psycho",	 "A wealthy New York investment banking executive hides his alternate psychopathic ego"	,2000	,"2000-04-14",	102,	4,	"R"),		
("The Dark Knight",	"The menace known as the Joker wreaks havoc on Gotham City",	2008	,"2008-07-18"	,152	,5,	"PG13");
	

INSERT INTO movies_genres (movie_id, genre_id ) VALUES (1,2),(2,3);

