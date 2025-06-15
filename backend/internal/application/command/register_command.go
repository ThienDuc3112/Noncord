package command

import "backend/internal/application/common"

type RegisterCommand struct {
	Name     string
	Email    string
	Password string
}

type RegisterCommandResult struct {
	Result *common.UserResult
}
