package request

import "net/http"

type NewServer struct {
}

func (r *NewServer) Bind(_ *http.Request) error {
	return validate.Struct(r)
}

type UpdateServer struct {
}

func (r *UpdateServer) Bind(_ *http.Request) error {
	return validate.Struct(r)
}
