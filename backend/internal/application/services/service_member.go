package services

import (
	"backend/internal/application/interfaces"
	"backend/internal/domain/repositories"
)

type MemberService struct {
	r repositories.MemberRepo
}

func NewMemberService(mr repositories.MemberRepo) interfaces.MembershipService {
	return nil
	// return &MemberService{
	// 	r: mr,
	// }
}
