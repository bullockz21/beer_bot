// internal/resource/user_resource.go
package resource

import (
	"github.com/bullockz21/beer_bot/internal/dto"
	"github.com/bullockz21/beer_bot/internal/entity"
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
