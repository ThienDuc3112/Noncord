package services

type InviteService interface {
	CreateInvite() error
	RevokeInvite() error
}
