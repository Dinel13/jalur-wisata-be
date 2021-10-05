package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	r := httprouter.New()

	r.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	r.ServeFiles("/assets/*filepath", http.Dir("images"))

	r.HandlerFunc(http.MethodPost, "/v1/signup", app.userSignup)
	r.HandlerFunc(http.MethodPost, "/v1/login", app.userLogin)

	r.HandlerFunc(http.MethodPost, "/v1/destiny", app.createDestinyHandler)
	r.HandlerFunc(http.MethodGet, "/v1/destiny/:id", app.getDestiny)

	return app.enableCors(r)
}
