package request

type LoginDTO struct {
	Email    string `json:"email" example:"john@gmail.com"`
	Password string `json:"password" example:"Test1234!"`
}
