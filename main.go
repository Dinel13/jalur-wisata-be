package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dinel13/wisata/models"
	_ "github.com/lib/pq"
)

const version = "0.0.1"

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
	jwt struct {
		secret string
	}
}

type appStatus struct {
	Version    string `json:"version"`
	Status     string `json:"status"`
	Enviroment string `json:"enviroment"`
}

type application struct {
	config config
	logger *log.Logger
	models models.Models
}

// var for connect to db
const (
	host     = "localhost"
	port     = "5432"
	user     = "din"
	password = "postgres"
	dbname   = "wisata"
)

func main() {
	var cfg config

	//set default config use flag in terminal
	flag.IntVar(&cfg.port, "port", 4000, "serverport listen on")
	flag.StringVar(&cfg.env, "env", "development", "Aplication enviroment (dev | prod) ")
	cfg.db.dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	flag.StringVar(&cfg.jwt.secret, "jwt-secret", "2dce505d96a53c5768052ee90f3df2055657518dad489160df9913f66042e160", "jwt secret key")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime) // create a new logger, 1 place to store, 2 prefix, 3 format

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err) // will stop the program
	}
	defer db.Close()

	app := &application{
		config: cfg,
		logger: logger,
		models: models.NewModel(db),
	}
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Starting server on port %d", app.config.port) //equal to cfg.port

	err = srv.ListenAndServe()

	if err != nil {
		log.Println(err)
	}
}

//retun a poiter connection to database
func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)

	if err != nil {
		return nil, err
	}
	return db, nil
}
