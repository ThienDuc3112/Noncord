package request

import "net/http"

type Register struct {
	Username string `json:"username" example:"tungsten_kitty" validate:"required,safe_username"`
	Email    string `json:"email" example:"tungstenkitty@gmail.com" validate:"required,email"`
	Password string `json:"password" example:"Password1@" validate:"required,min=8,max=72"`
}

func (r *Register) Bind(_ *http.Request) error {
	return validate.Struct(r)
}

type Login struct {
	Username string `json:"username" example:"tungsten_kitty" validate:"required"`
	Password string `json:"password" example:"Password1@" validate:"required"`
}

func (r *Login) Bind(_ *http.Request) error {
	return validate.Struct(r)
}

type Refresh struct {
	RefreshToken string `json:"refreshToken" example:"9Vz6ayzM0scQSIXHtYVbKcDeF1aa0aLs" validate:"required,len=32"`
}

func (r *Refresh) Bind(_ *http.Request) error {
	return validate.Struct(r)
}
