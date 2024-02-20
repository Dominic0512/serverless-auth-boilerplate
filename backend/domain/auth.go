package domain

type OAuthSignUpInput struct {
	Code  string
	State string
}

type OAuthSignInInput struct {
	Code  string
	State string
}
