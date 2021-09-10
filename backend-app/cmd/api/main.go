package main

import (
	"flag"
	"fmt"
	"net/http"
	"log"
	"os"
	"time"
	"database/sql"
	"context"
	_ "../../mysql-1.2"
	"../../models"
)


const version = "1.0.0"

type config struct {
	port int
	env string
	db struct {
		dsn string 
	}
}

type AppStatus struct {
	Status string `json:"status"`
	Environment string `json:"env"`
	Version string `json:"version"`
}

type application struct{
	config config
	logger *log.Logger	
	models models.Models
	
}

func main(){
	var cfg config
	
	flag.IntVar(&cfg.port, "port", 80, "server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "application environment(development|production)")
	flag.StringVar(&cfg.db.dsn,"dsn", "root@tcp(127.0.0.1:3306)/go_movies_db?parseTime=true", "Mysql address connection")
	flag.Parse()
	
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	
	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	
	defer db.Close()
	
	app := &application {
		config : cfg,
		logger : logger,
		models : models.NewModels(db),
	}
	
	
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.port),
		Handler : app.routes(),
		IdleTimeout : time.Minute,
		ReadTimeout : 10 * time.Second,
		WriteTimeout : 30 * time.Second,
	}
	
	logger.Println("Starting server on port", cfg.port)
	
	err = srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}


func openDB(cfg config) (*sql.DB, error){
	db,err := sql.Open("mysql", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	
	return db, nil
}