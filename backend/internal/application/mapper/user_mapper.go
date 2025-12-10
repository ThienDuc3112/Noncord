package mapper

import (
	"backend/internal/application/common"
	"backend/internal/domain/entities"

	"github.com/google/uuid"
)

func NewUserResultFromUserEntity(user *entities.User) *common.UserResult {
	if user == nil {
		return nil
	}

	return &common.UserResult{
		Id:          uuid.UUID(user.Id),
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		Username:    user.Username,
		DisplayName: user.DisplayName,
		AboutMe:     user.AboutMe,
		Email:       user.Email,
		Disabled:    user.Disabled,
		Verified:    user.Verified,
		AvatarUrl:   user.AvatarUrl,
		BannerUrl:   user.BannerUrl,
		Flags:       uint16(user.Flags),
	}
}
