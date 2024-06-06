package model

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64   `json:"id" db:"id"`
	UUID      string  `json:"uuid" db:"uuid"`
	Name      string  `json:"name" db:"name"`
	Email     string  `json:"email" db:"email"`
	Gender    string  `json:"gender" db:"gender"`         // "male" or "female"
	BirthDate string  `json:"birth_date" db:"birth_date"` // YYYY-MM-DD
	Password  string  `json:"password" db:"password"`
	CreatedAt string  `json:"created_at" db:"created_at"`
	UpdatedAt string  `json:"updated_at" db:"updated_at"`
	DeletedAt *string `json:"deleted_at" db:"deleted_at"`
}

const (
	ERROR_CREATE   = "CreateUser failed: %v"
	ERROR_GET_USER = "GetUserByEmail failed: %v"
)

func CreateUser(db *sqlx.DB, userData User) (*User, error) {
	query := `INSERT INTO users (uuid, name, email, gender, birth_date, password) VALUES (?,?,?,?,?,?)`

	tx, err := db.BeginTx(context.TODO(), nil)
	if err != nil {
		return nil, fmt.Errorf(ERROR_CREATE, err)
	}

	userData.UUID = uuid.New().String()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(userData.Password))
	if err != nil {
		return nil, fmt.Errorf(ERROR_CREATE, err)
	}

	_, err = tx.Exec(query, userData.UUID, userData.Name, userData.Email, userData.Gender, userData.BirthDate, hashedPassword)
	if err != nil {
		return nil, fmt.Errorf(ERROR_CREATE, err)
	}

	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
		return nil, fmt.Errorf(ERROR_CREATE, err)
	}

	result, err := GetUserByEmail(db, userData.Email)
	if err != nil {
		return nil, fmt.Errorf(ERROR_CREATE, err)
	}

	return result, nil
}

func GetUserByEmail(db *sqlx.DB, email string) (*User, error) {
	var user User
	query := `SELECT id, uuid, name, email, gender, birth_date, password, created_at, updated_at, deleted_at FROM users WHERE email = ? AND deleted_at IS NULL LIMIT 1`

	err := db.Get(&user, query, email)
	if err != nil {
		return nil, fmt.Errorf(ERROR_GET_USER, err)
	}

	return &user, nil
}
