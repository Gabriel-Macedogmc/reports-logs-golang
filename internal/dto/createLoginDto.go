package dto

type CreateLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
