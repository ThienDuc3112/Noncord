package command

type LoginCommand struct {
	Username  string
	Password  string
	UserAgent string
}

type LoginCommandResult struct {
	AccessToken  string
	RefreshToken string
}
