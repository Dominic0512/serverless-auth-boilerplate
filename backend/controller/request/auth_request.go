package request

type SignUpRequest struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,gte=6,eqfield=Password"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=ConfirmPassword"`
}
