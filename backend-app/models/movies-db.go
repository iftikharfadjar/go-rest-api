package models

import (
	"database/sql"
	"context"
	"time"
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
		genres[mg.ID]=mg.Genre.GenreName
	}
	
	movie.MovieGenre = genres;
	
	return &movie, nil
}

func (m *DBModel) All() ([]*Movie, error){
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	rows, err := m.DB.QueryContext(ctx, "SELECT * FROM movies ORDER BY title")
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
			genres[mg.ID]=mg.Genre.GenreName
			
		}
		genreRows.Close()

		movie.MovieGenre = genres;
		movies = append(movies, &movie)
		}
		return movies, nil
}