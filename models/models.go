package models

import (
	"database/sql"
	"time"
)

// models is wrapper for databas
type Models struct {
	DB DBModel
}

// NewMOdel retuiurn models with db pool
func NewModel(db *sql.DB) Models {
	return Models{
		DB: DBModel{
			DB: db,
		},
	}
}

// User is model for user
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Destiny is model for destiny
type Destiny struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Rating      float64   `json:"rating"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

// Category is model for category
type Category struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// DestinyCategory is model for destiny category
// it use to join destiny and category
type DestinyCategory struct {
	ID         int       `json:"id"`
	DestinyID  int       `json:"destiny_id"`
	CategoryID int       `json:"category_id"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}
