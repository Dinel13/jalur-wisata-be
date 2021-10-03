package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dinel13/wisata/models"
	"github.com/pascaldekloe/jwt"
	"golang.org/x/crypto/bcrypt"
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

	// hash the password
	passwrod, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 14)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	user.Name = name
	user.Email = payload.Email
	user.Password = string(passwrod)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	newUser, err := app.models.DB.CreateUser(user)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	var claims jwt.Claims
	claims.Subject = fmt.Sprint(newUser.ID)
	claims.Issued = jwt.NewNumericTime(time.Now())
	claims.NotBefore = jwt.NewNumericTime(time.Now())
	claims.Expires = jwt.NewNumericTime(time.Now().Add(time.Hour * 24))
	claims.Issuer = "test"
	claims.Audiences = []string{"test"}

	jwtBytes, err := claims.HMACSign(jwt.HS256, []byte(app.config.jwt.secret))
	if err != nil {
		app.errorJSON(w, errors.New("error signin"))
		return
	}

	// return the user created with the token
	type response struct {
		Token string `json:"token"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	resp := response{
		Token: string(jwtBytes),
		Name:  newUser.Name,
		Email: newUser.Email,
	}

	// return the user
	app.writeJSON(w, http.StatusOK, resp, "user")
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
