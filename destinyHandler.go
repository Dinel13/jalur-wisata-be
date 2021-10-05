package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/dinel13/wisata/models"
	"github.com/julienschmidt/httprouter"
)

type DestinyResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
	Images      string  `json:"images"`
	Category    string  `json:"category"`
}

// createDestinyHandler is a handler for the createDestiny function.
func (app *application) createDestinyHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseMultipartForm(1024); err != nil {
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	name := r.FormValue("name")
	description := r.FormValue("description")
	category := r.FormValue("category")
	rating, err := strconv.ParseFloat((r.FormValue("rating")), 64)
	if err != nil {
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	fmt.Println(rating)

	uploadedImage, header, err := r.FormFile("image")
	if err != nil {
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	defer uploadedImage.Close()

	dir, err := os.Getwd()
	if err != nil {
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	filename := header.Filename
	if name != "" {
		filename = fmt.Sprintf("%s%s", name, filepath.Ext(header.Filename))
	}

	fileLocation := filepath.Join(dir, "images", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedImage); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var destiny models.Destiny

	destiny.Name = name
	destiny.Description = description
	destiny.Rating = rating
	destiny.Image = filename
	destiny.CreatedAt = time.Now()
	destiny.UpdatedAt = time.Now()
	destiny.Category = category

	newDestiny, err := app.models.DB.CreateDestiny(destiny)
	if err != nil {
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	//return destiny as destinyResponse
	destinyResponse := DestinyResponse{
		ID:          newDestiny.ID,
		Name:        newDestiny.Name,
		Description: newDestiny.Description,
		Rating:      newDestiny.Rating,
		Images:      newDestiny.Image,
		Category:    newDestiny.Category,
	}

	// return the user
	app.writeJSON(w, http.StatusOK, destinyResponse, "destiny")
}

// getDestiny is handler for get one destyny by id
func (app *application) getDestiny(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	fmt.Println(id)
	if err != nil {
		app.errorJSON(w, err, http.StatusInternalServerError)
	}

	destiny, err := app.models.DB.GetDestiny(id)
	if err != nil {
		app.errorJSON(w, err, http.StatusInternalServerError)
	}

	app.writeJSON(w, http.StatusOK, destiny, "destiny")
}
