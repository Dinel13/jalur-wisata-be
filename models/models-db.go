package models

import (
	"context"
	"database/sql"
	"fmt"
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

// GetDestiny return one destiny by id
func (db *DBModel) GetDestiny(id int) (*Destiny, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT * FROM destinies WHERE id = $1`

	row := db.DB.QueryRowContext(ctx, query, id)

	destiny := &Destiny{}

	err := row.Scan(
		&destiny.ID,
		&destiny.Name,
		&destiny.Description,
		&destiny.Rating,
		&destiny.Image,
		&destiny.CreatedAt,
		&destiny.UpdatedAt,
		&destiny.Category,
	)
	// cek if row is empty
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return destiny, nil
}

// GetAllDestinies returns all destinies
func (db *DBModel) GetAllDestinies() ([]*Destiny, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT * FROM destinies`
	rows, err := db.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	destinies := []*Destiny{}
	for rows.Next() {
		destiny := &Destiny{}
		err := rows.Scan(
			&destiny.ID,
			&destiny.Name,
			&destiny.Description,
			&destiny.Rating,
			&destiny.Image,
			&destiny.CreatedAt,
			&destiny.UpdatedAt,
			&destiny.Category,
		)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		destinies = append(destinies, destiny)
	}

	return destinies, nil
}

// UpdateDestiny updates a destiny
func (db *DBModel) UpdateDestiny(id int, destiny Destiny) (*Destiny, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	sqlStatement := `UPDATE destinies SET name = $1, description = $2, rating = $3, image = $4, updated_at = $5, category = $6 WHERE id = $7 RETURNING *`

	row := db.DB.QueryRowContext(ctx, sqlStatement,
		destiny.Name,
		destiny.Description,
		destiny.Rating,
		destiny.Image,
		destiny.UpdatedAt,
		destiny.Category,
		id,
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

	if err == sql.ErrNoRows {
		fmt.Println("Destiny not found", id)
		return nil, nil
	}

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &newDestiny, nil
}

// DeleteDestiny deletes a destiny
func (db *DBModel) DeleteDestiny(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	sqlStatement := `DELETE FROM destinies WHERE id = $1`

	_, err := db.DB.ExecContext(ctx, sqlStatement, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// GetDestiniesByCategory returns all destinies by category
func (db *DBModel) GetDestiniesByCategory(category string) ([]*Destiny, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT * FROM destinies WHERE category = $1`
	rows, err := db.DB.QueryContext(ctx, query, category)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	destinies := []*Destiny{}
	for rows.Next() {
		destiny := &Destiny{}
		err := rows.Scan(
			&destiny.ID,
			&destiny.Name,
			&destiny.Description,
			&destiny.Rating,
			&destiny.Image,
			&destiny.CreatedAt,
			&destiny.UpdatedAt,
			&destiny.Category,
		)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		destinies = append(destinies, destiny)
	}

	return destinies, nil
}

// GetPOpularDestinies returns most popular destinies
func (db *DBModel) GetPopularDestinies() ([]*Destiny, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT * FROM destinies ORDER BY rating DESC LIMIT 8`
	rows, err := db.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	destinies := []*Destiny{}
	for rows.Next() {
		destiny := &Destiny{}
		err := rows.Scan(
			&destiny.ID,
			&destiny.Name,
			&destiny.Description,
			&destiny.Rating,
			&destiny.Image,
			&destiny.CreatedAt,
			&destiny.UpdatedAt,
			&destiny.Category,
		)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		destinies = append(destinies, destiny)
	}

	return destinies, nil
}

// GetLatestDestinies returns latest destinies
func (db *DBModel) GetLatestDestinies() ([]*Destiny, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT * FROM destinies ORDER BY created_at DESC LIMIT 8`
	rows, err := db.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	destinies := []*Destiny{}
	for rows.Next() {
		destiny := &Destiny{}
		err := rows.Scan(
			&destiny.ID,
			&destiny.Name,
			&destiny.Description,
			&destiny.Rating,
			&destiny.Image,
			&destiny.CreatedAt,
			&destiny.UpdatedAt,
			&destiny.Category,
		)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		destinies = append(destinies, destiny)
	}

	return destinies, nil
}
