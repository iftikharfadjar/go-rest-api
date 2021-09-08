package main

import (
	"net/http"
	"../../httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()
	
	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)
	
	return router
}