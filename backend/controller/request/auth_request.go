package request

type SignUpRequest struct {
	Code string `json:"code" validate:"required"`
}

type SignInRequest struct {
	Code string `json:"code" validate:"required"`
}
