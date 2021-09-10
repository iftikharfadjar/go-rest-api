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
	
	
	// query := `SELECT id,title,description,year,release_date, rating,runtime, mpaa_rating,created_at, updated_at FROM movies where id = 1;`
	
	row := m.DB.QueryRowContext(ctx, "SELECT id,title,description,year,release_date, rating,runtime, mpaa_rating,created_at, updated_at FROM movies where id = ?", id)
	
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
	return &movie, nil
}

func (m *DBModel) All(id int) ([]*Movie, error){
	return nil, nil
}