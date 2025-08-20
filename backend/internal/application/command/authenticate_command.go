package command

import "backend/internal/application/common"

type AuthenticateCommand struct {
	AccessToken string
}

type AuthenticateCommandResult struct {
	User *common.UserResult
}
