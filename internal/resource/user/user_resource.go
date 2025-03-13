package resource

import (
	dto "github.com/bullockz21/beer_bot/internal/dto/user"
	entity "github.com/bullockz21/beer_bot/internal/entity/user"
)

type UserResource struct{}

func NewUserResource() *UserResource {
	return &UserResource{}
}

func (r *UserResource) ToResponse(user *entity.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:         user.ID,
		Username:   user.Username,
		FirstName:  user.FirstName,
		Registered: user.ID != 0,
	}
}
