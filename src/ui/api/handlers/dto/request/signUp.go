package request

type SignUpDTO struct {
	Email     string `json:"email" example:"jose@gmail.com"`
	FirstName string `json:"first_name" example:"José"`
	LastName  string `json:"last_name" example:"Santos"`
	Password  string `json:"password" example:"Test1234!"`
}
