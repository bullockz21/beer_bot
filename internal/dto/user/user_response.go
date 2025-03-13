package dto

type UserResponse struct {
	ID         uint   `json:"id"`
	Username   string `json:"username"`
	FirstName  string `json:"first_name"`
	Registered bool   `json:"registered"`
}
