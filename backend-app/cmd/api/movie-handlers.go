package main

import (
	"net/http"
	"../../httprouter"
	"../../models"
	"strconv"
	"errors"
	"time"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request){
	params := httprouter.ParamsFromContext(r.Context());
	
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil{
		app.logger.Print(errors.New("Invalid id Parameter"))
	}
	
	movie := models.Movie {
		ID: id,
		Title: "Some Movie",
		Description: "Some Description",
		Year: 2021,
		ReleaseDate: time.Date(2021,01,01,01,0,0,0,time.Local),
		Runtime: 100,
		Rating: 5,
		MPAARating: "PG_13",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	
	err = app.writeJSON(w, http.StatusOK, "movie", movie)
	
	
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request){
	
}