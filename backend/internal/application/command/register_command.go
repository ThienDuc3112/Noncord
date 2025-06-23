package command

import "backend/internal/application/common"

type RegisterCommand struct {
	Username string
	Email    string
	Password string
}

type RegisterCommandResult struct {
	Result *common.UserResult
}
