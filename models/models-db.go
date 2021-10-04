package models

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

// RegisterUser inserts a new user into the database
func (db *DBModel) CreateUser(user User) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// we use queryrow becouse execContext cant return the user inserted

	// stmt := `INSERT INTO users (name, email, password, created_at, updated_at ) VALUES ($1, $2, $3, $4, $5)`
	// result, err := db.DB.ExecContext(ctx, stmt,
	// 	user.Name,
	// 	user.Email,
	// 	user.Password,
	// 	user.CreatedAt,
	// 	user.UpdatedAt,
	// )

	sqlStatement := `INSERT INTO users (name, email, password, created_at, updated_at )
							VALUES ($1, $2, $3, $4, $5)
							RETURNING *`

	row := db.DB.QueryRowContext(ctx, sqlStatement,
		user.Name,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	)

	var newUser User
	err := row.Scan(&newUser.ID, &newUser.Name, &newUser.Email, &newUser.Password, &newUser.CreatedAt, &newUser.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &newUser, nil
}

// GetUserByEmail returns a user by email
func (db *DBModel) GetUserByEmail(email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	log.Println(email)

	query := `SELECT * FROM users WHERE email = $1`
	row := db.DB.QueryRowContext(ctx, query, email)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		log.Println(err, "sadsa")
		return nil, err
	}

	return &user, nil
}

// GetUserByID returns a user by id
func (db *DBModel) GetUserByID(id int) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `SELECT * FROM users WHERE id = $1`
	row := db.DB.QueryRowContext(ctx, stmt, id)

	user := &User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return user, nil
}

// GetAllUsers returns all users
func (db *DBModel) GetAllUsers() ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `SELECT * FROM users`
	rows, err := db.DB.QueryContext(ctx, stmt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	users := []*User{}
	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// CreateDestiny inserts a new destiny into the database
func (db *DBModel) CreateDestiny(destiny Destiny) (*Destiny, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	sqlStatement := `INSERT INTO destinies (name, description, rating, image, created_at, updated_at, category)
							VALUES ($1, $2, $3, $4, $5, $6, $7)
							RETURNING *`

	row := db.DB.QueryRowContext(ctx, sqlStatement,
		destiny.Name,
		destiny.Description,
		destiny.Rating,
		destiny.Image,
		destiny.CreatedAt,
		destiny.UpdatedAt,
		destiny.Category,
	)

	var newDestiny Destiny
	err := row.Scan(
		&newDestiny.ID,
		&newDestiny.Name,
		&newDestiny.Description,
		&newDestiny.Rating,
		&newDestiny.Image,
		&newDestiny.CreatedAt,
		&newDestiny.UpdatedAt,
		&newDestiny.Category,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &newDestiny, nil
}
