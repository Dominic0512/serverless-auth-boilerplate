package response

type GenerateAuthURLResponse struct {
	Url *string `json:"url"`
}

type TokenResponse struct {
	Token      string `json:"token"`
	Token_type string `json:"token_type"`
}
