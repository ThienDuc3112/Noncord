package postgres

import (
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
		DeletedAt:   nil,
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
	if user.DeletedAt.Valid {
		res.DeletedAt = &user.DeletedAt.Time
	}

	return res
}
