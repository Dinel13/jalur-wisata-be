package main

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) wrap(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), "allParams", ps)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (app *application) routes() http.Handler {
	r := httprouter.New()

	secure := alice.New(app.checkToken)

	r.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	r.ServeFiles("/assets/*filepath", http.Dir("images"))

	r.HandlerFunc(http.MethodPost, "/v1/signup", app.userSignup)
	r.HandlerFunc(http.MethodPost, "/v1/login", app.userLogin)

	r.POST("/v1/destiny", app.wrap(secure.ThenFunc(app.createDestinyHandler)))
	// r.POST("/v1/destiny", app.wrap(secure.ThenFunc(app.createDestiny)))
	r.POST("/v1/destiny/:id", app.wrap(secure.ThenFunc(app.updateDestiny)))

	r.HandlerFunc(http.MethodGet, "/v1/destiny/:id", app.getDestiny)
	r.HandlerFunc(http.MethodGet, "/v1/destinies", app.getAllDestiny)
	r.HandlerFunc(http.MethodGet, "/v1/pop-destinies", app.popularDestiny)
	r.HandlerFunc(http.MethodGet, "/v1/latest-destinies", app.latestDestiny)

	// r.HandlerFunc(http.MethodDelete, "/v1/destiny/:id", app.deleteDestiny)

	return app.enableCors(r)
}
