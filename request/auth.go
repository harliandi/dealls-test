package request

type SignupRequest struct {
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Gender    string `json:"gender" validate:"required,oneof=male female prefer_not_to"` // "male" or "female"
	BirthDate string `json:"birth_date" validate:"birth-date"`                           // YYYY-MM-DD
	Password  string `json:"password" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
