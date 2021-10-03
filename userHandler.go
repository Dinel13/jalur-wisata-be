package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dinel13/wisata/models"
)

type userPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// signupHadler handles the signup request
func (app *application) signupHandler(w http.ResponseWriter, r *http.Request) {
	// get the email and password from the body request
	var payload userPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Printf("Error decoding payload: %v", err)
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	// check if the user already exists
	existUser, err := app.models.DB.GetUserByEmail(payload.Email)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	if existUser != nil {
		log.Printf("User already exists")
		app.errorJSON(w, errors.New("user already exist"), http.StatusForbidden)
		return
	}

	// create the user
	var user models.User

	// get name by splitting email
	name := payload.Email
	if i := strings.LastIndex(payload.Email, "@"); i > 0 {
		name = payload.Email[:i]
	}
	user.Name = name
	user.Email = payload.Email
	user.Password = payload.Password
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	newUser, err := app.models.DB.CreateUser(user)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	// return the user
	app.writeJSON(w, http.StatusOK, newUser, "user")
}
