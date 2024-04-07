package authenticator

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/Dominic0512/serverless-auth-boilerplate/infra/config"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type Auth0Authenticator struct {
	*oidc.Provider
	*oidc.IDTokenVerifier
	oauth2.Config
}

func generateRandomState() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(bytes)

	return state, nil
}

func (auth *Auth0Authenticator) GenerateAuthCodeURL() (string, error) {
	randState, err := generateRandomState()
	if err != nil {
		return "", err
	}

	return auth.AuthCodeURL(randState), nil
}

func (auth *Auth0Authenticator) ExchangeMetaDataByCode(code string) (*AuthMetaData, error) {
	oauth2Token, err := auth.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		return nil, fmt.Errorf("can not extract the id token from oauth2 token")
	}
	fmt.Println(rawIDToken)
	idToken, err := auth.Verify(context.Background(), rawIDToken)
	if err != nil {
		return nil, err
	}

	var claims struct {
		Email    string `json:"email"`
		Verified bool   `json:"email_verified"`
		Picture  string `json:"picture"`
		Sub      string `json:"sub"`
	}

	if err := idToken.Claims(&claims); err != nil {
		return nil, err
	}

	return &AuthMetaData{
		AccessToken:   oauth2Token.AccessToken,
		Email:         claims.Email,
		EmailVerified: claims.Verified,
		Picture:       claims.Picture,
		Sub:           claims.Sub,
	}, nil
}

func NewAuth0Authenticator(config *config.Config) (*Auth0Authenticator, error) {
	domain := config.Auth0Domain
	clientID := config.Auth0ClientID
	clientSecret := config.Auth0ClientSecret
	callbackURL := config.Auth0CallbackURL

	provider, err := oidc.NewProvider(
		context.Background(),
		fmt.Sprintf("https://%s/", domain),
	)
	if err != nil {
		return nil, err
	}

	verifier := provider.Verifier(&oidc.Config{ClientID: config.Auth0ClientID})

	oConfig := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  callbackURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	return &Auth0Authenticator{
		Provider:        provider,
		IDTokenVerifier: verifier,
		Config:          oConfig,
	}, nil
}
