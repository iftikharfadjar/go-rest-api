package main

import (
	"net/http"
	"../../httprouter"
	"strconv"
	"errors"
	"../../models"
	"time"
	"encoding/json"
)

	type jsonResp struct{
		OK bool `json:"ok"`
		Message string `json:"message"`
	}

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request){
	params := httprouter.ParamsFromContext(r.Context());
	
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil{
		app.logger.Print(errors.New("Invalid id Parameter"))
		app.errorJSON(w, err)
		return
	}
	
	movie, err := app.models.DB.Get(id)
	
	if err != nil {
		app.logger.Print(errors.New("Invalid id Parameter"))
		app.errorJSON(w, err)
		return
	}

	
	err = app.writeJSON(w, http.StatusOK, "movie", movie)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request){
	movies , err := app.models.DB.All()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	
	err = app.writeJSON(w, http.StatusOK, "movies", movies)
		if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getAllGenres(w http.ResponseWriter, r *http.Request){
	genres, err := app.models.DB.GenresAll()
	
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	
	err = app.writeJSON(w, http.StatusOK, "genres", genres)
		if err != nil {
		app.errorJSON(w, err)
		return
	}
	
}

func (app *application) getAllMoviesByGenre(w http.ResponseWriter, r *http.Request){
	params := httprouter.ParamsFromContext(r.Context());
	
	genreID, err := strconv.Atoi(params.ByName("genre_id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	
	movies , err := app.models.DB.All(genreID)
	
	err = app.writeJSON(w, http.StatusOK, "movies", movies)
		if err != nil {
			app.errorJSON(w, err)
		return
	}
}

func (app *application) moviesByGenre(w http.ResponseWriter, r *http.Request){
}

func (app *application) deleteMovie(w http.ResponseWriter, r *http.Request){
}

type MoviePayload struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Year string  `json:"year"`
	ReleaseDate string `json:"release_date"`
	Runtime string `json:"runtime"`
	Rating string `json:"rating"`
	MPAARating string `json:"mpaa_rating"`
}

func (app *application) editMovie(w http.ResponseWriter, r *http.Request){
	// var movie models.Movie
	var payload MoviePayload
	
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err!= nil {
		app.errorJSON(w,err)
		return
	}
	
	var movie models.Movie
	
	movie.ID, _ 	= strconv.Atoi(payload.ID)
	movie.Title = payload.Title
	movie.Description = payload.Description
	movie.ReleaseDate,_ = time.Parse("2005-01-02", payload.ReleaseDate)
	movie.Year = movie.ReleaseDate.Year()
	movie.Runtime , _ = strconv.Atoi(payload.Runtime)
	movie.Rating , _ = strconv.Atoi(payload.Rating)
	movie.MPAARating = payload.MPAARating
	movie.CreatedAt = time.Now()
	movie.UpdatedAt = time.Now()
	
	err = app.models.DB.InsertMovie(movie)
	if err != nil {
			app.errorJSON(w, err)
		return
	}
	
	
	ok := jsonResp{
		OK : true,
	}
	
	err = app.writeJSON(w, http.StatusOK, "response", ok)
		if err != nil {
			app.errorJSON(w, err)
		return
	}
	
}

func (app *application) updateMovie(w http.ResponseWriter, r *http.Request){
}

func (app *application) searchMovies(w http.ResponseWriter, r *http.Request){
}