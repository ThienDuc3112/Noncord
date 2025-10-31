package request

import "net/http"

type UpdateInvitation struct {
	BypassApproval *bool  `json:"bypassApproval" validate:"required_without_all=JoinLimit"`
	JoinLimit      *int32 `json:"joinLimit" validate:"required_without_all=BypassApproval"`
}

func (r *UpdateInvitation) Bind(_ *http.Request) error {
	return validate.Struct(r)
}
