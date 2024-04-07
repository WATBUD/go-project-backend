package userRequest

type CreateUserRequest struct {
	Account  string `json:"account"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
