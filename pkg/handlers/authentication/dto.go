package authentication

type SignInDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponseDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
