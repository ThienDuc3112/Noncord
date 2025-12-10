package postgres

import (
	"backend/internal/application/common"
	"backend/internal/domain/entities"
	"backend/internal/infra/db/postgres/gen"
)

func fromDbUser(user *gen.User) *entities.User {
	if user == nil {
		return nil
	}

	res := &entities.User{
		Id:          entities.UserId(user.ID),
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		DeletedAt:   user.DeletedAt,
		Username:    user.Username,
		DisplayName: user.DisplayName,
		AboutMe:     user.AboutMe,
		Email:       user.Email,
		Password:    "",
		Disabled:    user.Disabled,
		Verified:    true,
		AvatarUrl:   user.AvatarUrl,
		BannerUrl:   user.BannerUrl,
		Flags:       entities.UserFlags(user.Flags),
	}
	if user.Password.Valid {
		res.Password = user.Password.String
	}

	return res
}

func toCommonUser(user gen.User) common.UserResult {
	return common.UserResult{
		Id:          user.ID,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		Username:    user.Username,
		DisplayName: user.DisplayName,
		AboutMe:     user.AboutMe,
		Email:       user.Email,
		Disabled:    user.Disabled,
		AvatarUrl:   user.AvatarUrl,
		BannerUrl:   user.BannerUrl,
		Flags:       uint16(user.Flags),
		Verified:    true,
	}
}
