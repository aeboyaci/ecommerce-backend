package authentication

type SignInDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
