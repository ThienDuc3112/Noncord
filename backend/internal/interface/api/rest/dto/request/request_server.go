package request

import "net/http"

type NewServer struct {
	Name string `json:"name" example:"My very good server" validate:"required,max=256"`
}

func (r *NewServer) Bind(_ *http.Request) error {
	return validate.Struct(r)
}

type UpdateServer struct {
}

func (r *UpdateServer) Bind(_ *http.Request) error {
	return validate.Struct(r)
}
