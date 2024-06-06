package dto

type CreateClient struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Document string `json:"document"`
}
