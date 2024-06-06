package model

import (
	"context"
	"dealls-test/request"
	"encoding/base64"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

func Signup(db *sqlx.DB, userData request.SignupRequest) error {
	_, err := CreateUser(db, User{
		Name:      userData.Name,
		Email:     userData.Email,
		Gender:    userData.Gender,
		BirthDate: userData.BirthDate,
		Password:  userData.Password,
	})
	if err != nil {
		return err
	}
	return nil
}

func Login(db *sqlx.DB, redis *redis.Client, email, password string) (string, error) {

	user, err := GetUserByEmail(db, email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token := base64.RawStdEncoding.EncodeToString([]byte(user.Email + "@" + user.Password + "@" + time.Now().Format(time.DateOnly)))

	_, err = redis.Set(context.TODO(), email, token, 24*time.Hour).Result()
	if err != nil {
		return "", err
	}
	return token, nil
}
