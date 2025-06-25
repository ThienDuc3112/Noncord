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
