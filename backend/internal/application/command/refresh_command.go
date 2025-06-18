package command

type RefreshCommand struct {
	RefreshToken string
}

type RefreshCommandResult struct {
	AccessToken  string
	RefreshToken string
}
