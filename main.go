package main

import (
	"encoding/json"
	"log"
	"net/http"

	"dealls-test/model"
	"dealls-test/request"
	"dealls-test/utils"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type DB struct {
	sql   *sqlx.DB
	redis *redis.Client
}

var Access *DB
var validate *validator.Validate

func init() {
	db, err := sqlx.Connect("mysql", "root:secret@tcp(localhost:3306)/dealls-db?parseTime=true")
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	Access = &DB{sql: db, redis: client}
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	validate = validator.New()
	validate.RegisterValidation("birth-date", utils.IsBirthDate)

	r.Post("/signup", func(w http.ResponseWriter, r *http.Request) {
		var signup request.SignupRequest

		err := json.NewDecoder(r.Body).Decode(&signup)
		if err != nil {
			utils.RespondWithJSON(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "invalid JSON")
			return
		}

		err = validate.Struct(signup)
		if err != nil {
			utils.RespondWithJSON(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
			return
		}

		err = model.Signup(Access.sql, signup)
		if err != nil {
			utils.RespondWithJSON(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, http.StatusText(http.StatusOK), "Registered successfully, please login using your email and password")
		return

	})

	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		var login request.LoginRequest

		err := json.NewDecoder(r.Body).Decode(&login)
		if err != nil {
			utils.RespondWithJSON(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "invalid JSON")
			return
		}

		err = validate.Struct(login)
		if err != nil {
			utils.RespondWithJSON(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
			return
		}

		token, err := model.Login(Access.sql, Access.redis, login.Email, login.Password)
		if err != nil {
			utils.RespondWithJSON(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, http.StatusText(http.StatusOK), struct{ Token string }{Token: token})
		return
	})
	log.Print("app started on port :3001")
	http.ListenAndServe(":3001", r)
}
