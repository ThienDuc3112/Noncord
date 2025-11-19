package command

import (
	"github.com/google/uuid"
)

type AuthenticateCommand struct {
	AccessToken string
}

type AuthenticateCommandResult struct {
	UserId *uuid.UUID
}
