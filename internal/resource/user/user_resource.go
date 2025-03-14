package user

import (
	dtoUser "github.com/bullockz21/beer_bot/internal/dto/user"
	entityUser "github.com/bullockz21/beer_bot/internal/entity/user"
)

type UserResource struct{}

func NewUserResource() *UserResource {
	return &UserResource{}
}

func (r *UserResource) ToResponse(user *entityUser.User) *dtoUser.UserResponse {
	return &dtoUser.UserResponse{
		ID:         user.ID,
		Username:   user.Username,
		FirstName:  user.FirstName,
		Registered: true,
	}
}
