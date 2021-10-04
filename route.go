package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	r := httprouter.New()

	r.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	r.HandlerFunc(http.MethodPost, "/v1/signup", app.userSignup)
	r.HandlerFunc(http.MethodPost, "/v1/login", app.userLogin)

	r.HandlerFunc(http.MethodPost, "/v1/destiny", app.createDestinyHandler)
	r.ServeFiles("/assets/*filepath", http.Dir("images"))

	return app.enableCors(r)
}
