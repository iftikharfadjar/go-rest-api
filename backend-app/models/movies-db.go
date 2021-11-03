package models

import (
	"database/sql"
	"context"
	"time"
	"fmt"
	"log"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) Get(id int) (*Movie, error){
	
	ctx, cancel := context.WithTimeout(context.Background() , 3* time.Second)
	defer cancel()
	
	
	row := m.DB.QueryRowContext(ctx, "SELECT * FROM movies where id = ?", id)
	
	var movie Movie
	
	err := row.Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.Year,
		&movie.ReleaseDate,
		&movie.Rating,
		&movie.Runtime,
		&movie.MPAARating,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)
	
	if err != nil {
		return nil, err
	}
	
	query := "select mg.id, mg.movie_id, mg.genre_id, g.genre_name from movies_genres mg LEFT JOIN genres g ON g.id=mg.genre_id WHERE mg.movie_id=?"
	rows, _ := m.DB.QueryContext(ctx, query,  id)
	defer rows.Close()
	
	// var genres []MovieGenre
	genres := make(map[int]string)
	for rows.Next(){
		var mg MovieGenre
		err = rows.Scan(
			&mg.ID,
			&mg.MovieID,
			&mg.GenreID,
			&mg.Genre.GenreName,
		)
		
		if err != nil {
			return nil, err
		}
		
		// genres = append(genres, mg)
		genres[mg.GenreID]=mg.Genre.GenreName
	}
	
	movie.MovieGenre = genres;
	
	return &movie, nil
}

func (m *DBModel) All(genre ...int) ([]*Movie, error){
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	where := ""
	if len(genre) > 0 {
		where = fmt.Sprintf("WHERE id in (SELECT movie_id FROM movies_genres where genre_id = %d)", genre[0])
	}
	
	query := fmt.Sprintf("SELECT * FROM movies %s ORDER BY title", where)
	
	
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var movies []*Movie
	

	
	for rows.Next(){
		var movie Movie
			err = rows.Scan(
				&movie.ID,
				&movie.Title,
				&movie.Description,
				&movie.Year,
				&movie.ReleaseDate,
				&movie.Rating,
				&movie.Runtime,
				&movie.MPAARating,
				&movie.CreatedAt,
				&movie.UpdatedAt,
		)
		
		if err != nil {
			return nil, err
		}
		
		genrequery := "select mg.id, mg.movie_id, mg.genre_id, g.genre_name from movies_genres mg LEFT JOIN genres g ON g.id=mg.genre_id WHERE mg.movie_id=?"
		
		genreRows, _ := m.DB.QueryContext(ctx, genrequery,  movie.ID)
		
		// var genres []MovieGenre
		genres := make(map[int]string)
		for genreRows.Next(){
			
			var mg MovieGenre
			err = genreRows.Scan(
				&mg.ID,
				&mg.MovieID,
				&mg.GenreID,
				&mg.Genre.GenreName,
			)

			if err != nil {
				return nil, err
			}

		// 	// genres = append(genres, mg)
			genres[mg.GenreID]=mg.Genre.GenreName
			
		}
		genreRows.Close()

		movie.MovieGenre = genres;
		movies = append(movies, &movie)
		}
		return movies, nil
}

func (m *DBModel) GenresAll() ([]*Genre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	
	defer cancel()
	
	query := "SELECT id, genre_name, created_at, updated_at from genres ORDER BY genre_name"
	
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	
	var genres []*Genre
	
	for rows.Next() {
		
		var g Genre
		err:= rows.Scan(
			&g.ID,
			&g.GenreName,
			&g.CreatedAt,
			&g.UpdatedAt,
		)
		
		if err != nil {
			return nil, err
		}
		
		genres = append(genres, &g)
		
	}
	
	return genres, nil
	
}

func (m *DBModel) InsertMovie(movie Movie) error {
	ctx, cancel := context.WithTimeout(context.Background() , 3*time.Second)
	
	defer cancel()
	
	query := "INSERT INTO movies (title, description, year, release_date, runtime, rating, mpaa_rating, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?,?)"
	
	
	_, err := m.DB.ExecContext(ctx, query,
						movie.Title,
						movie.Description,
						movie.Year,
						movie.ReleaseDate,
						movie.Runtime,
						movie.Rating,
						movie.MPAARating,
						movie.CreatedAt,
						movie.UpdatedAt,
						)
	
	if err != nil {
		log.Println(err)
		return err
	}
	
	return nil
}

func (m *DBModel) UpdateMovie(movie Movie) error {
	ctx, cancel := context.WithTimeout(context.Background() , 3*time.Second)
	
	defer cancel()
	
	query := "UPDATE movies SET title=?, description=?, year=?, release_date=?, runtime=?, rating=?, mpaa_rating=?, created_at=?, updated_at=? WHERE id=?"
	
	
	_, err := m.DB.ExecContext(ctx, query,
						movie.Title,
						movie.Description,
						movie.Year,
						movie.ReleaseDate,
						movie.Runtime,
						movie.Rating,
						movie.MPAARating,
						movie.CreatedAt,
						movie.UpdatedAt,
					 	movie.ID,
						)
	
	if err != nil {
		log.Println(movie.ReleaseDate)
		log.Println(err)
		return err
	}
	
	return nil
}

func (m *DBModel) DeleteMovie (id int) error {
	ctx, cancel := context.WithTimeout(context.Background() , 3*time.Second)
	
	defer cancel()
	
	query := "DELETE FROM movies WHERE id=?"
	
	_, err := m.DB.ExecContext(ctx, query,id)
	
	if err != nil {
		log.Println(err)
		return err
	}
	
	return nil
	
}