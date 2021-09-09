package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) writeJSON(w http.ResponseWriter, status int,  wrap string,data interface{}) error{
	
	wrapper := make(map[string]interface{})
	
	wrapper[wrap] = data
	
	js,err := json.Marshal(wrapper)
	if err!= nil {
		return err
	}
	
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(status)
	w.Write(js)
	
	return nil
	
}

func (app *application) errorJSON(w http.ResponseWriter, err error){
	type jsonError struct {
		Message string `json:"message"`
	}
	
	theErr := jsonError{
		Message : err.Error(),
	}
	
	app.writeJSON(w, http.StatusBadRequest, "error" , theErr)
}