package main

import (
	"net/http"
	"../../httprouter"
	"strconv"
	"errors"
)

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

func (app *application) deleteMovie(w http.ResponseWriter, r *http.Request){
}

func (app *application) insertMovie(w http.ResponseWriter, r *http.Request){
}

func (app *application) updateMovie(w http.ResponseWriter, r *http.Request){
}

func (app *application) searchMovies(w http.ResponseWriter, r *http.Request){
}