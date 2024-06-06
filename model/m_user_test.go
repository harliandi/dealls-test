package model

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func mockDB(t *testing.T) *sqlx.DB {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dbx := sqlx.NewDb(db, "mysql")
	return dbx
}

func TestGetUserByEmailNotFoundEmail(t *testing.T) {
	// Test case: GetUserByEmail returns nil and an error when user not found
	db := mockDB(t)

	notFoundEmail := "notfound@email.com"

	user, err := GetUserByEmail(db, notFoundEmail)

	assert.Nil(t, user)
	assert.Error(t, err)
}

func TestGetUserByEmailAndPasswordSuccess(t *testing.T) {
	// Test case: GetUserByEmail returns user and nil error when user found
	db := mockDB(t)

	foundEmail := "found@email.com"

	user := User{
		ID:        1,
		UUID:      "uuid",
		Name:      "name",
		Email:     "found@email.com",
		Gender:    "gender",
		BirthDate: "birth_date",
		Password:  "password",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
		DeletedAt: nil,
	}
	result, err := GetUserByEmail(db, foundEmail)

	assert.Equal(t, user, result)
	assert.NoError(t, err)
}

func TestCreateUserSuccess(t *testing.T) {
	// Test case: CreateUser returns user data and nil error when create user success
	access := mockDB(t)

	user := User{
		ID:        1,
		UUID:      "uuid",
		Name:      "name",
		Email:     "email",
		Gender:    "gender",
		BirthDate: "birth_date",
		Password:  "pwd",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
		DeletedAt: nil,
	}

	result, err := CreateUser(access, user)

	assert.Equal(t, result, user)
	assert.NoError(t, err)

}

func TestCreateUserFailed(t *testing.T) {
	// Test case: CreateUser returns nil and an error when create user failed
	access := mockDB(t)

	user := User{
		ID:        1,
		UUID:      "uuid",
		Name:      "name",
		Email:     "email",
		Gender:    "gender",
		BirthDate: "birth_date",
		Password:  "pwd",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
		DeletedAt: nil,
	}

	result, err := CreateUser(access, user)

	assert.Nil(t, result)
	assert.Error(t, err)

}
