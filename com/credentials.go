package com

import (
	"context"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// OAuthCredentials is the interface for the different types of credentials
type OAuthCredentials interface {
	GetTokenSource(ctx context.Context, tokenURL string) (oauth2.TokenSource, error)
}

// PasswordCredentials are the credentials for the password grant type
type PasswordCredentials struct {
	Username string
	Password string
	Scopes   []string
}

// NewPasswordCredentials creates a new PasswordCredentials struct with the given parameters
func NewPasswordCredentials(username, password string, scopes []string) PasswordCredentials {
	return PasswordCredentials{
		Username: username,
		Password: password,
		Scopes:   scopes,
	}
}

// GetTokenSource returns the token source for the password grant type
func (c PasswordCredentials) GetTokenSource(ctx context.Context, tokenURL string) (oauth2.TokenSource, error) {
	oauthConf := &oauth2.Config{
		ClientID: "administration",
		Scopes:   c.Scopes,
		Endpoint: oauth2.Endpoint{
			TokenURL: tokenURL,
		},
	}

	token, err := oauthConf.PasswordCredentialsToken(ctx, c.Username, c.Password)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get token for %s", c.Username)
	}
	return oauth2.StaticTokenSource(token), nil
}

// IntegrationCredentials are the credentials for the authorization using a API key
type IntegrationCredentials struct {
	ClientId     string
	ClientSecret string
	Scopes       []string
}

// NewIntegrationCredentials creates a new IntegrationCredentials struct with the given parameters
func NewIntegrationCredentials(clientId, clientSecret string, scopes []string) IntegrationCredentials {
	return IntegrationCredentials{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		Scopes:       scopes,
	}
}

// GetTokenSource returns the token source for the client credentials grant type
func (c IntegrationCredentials) GetTokenSource(ctx context.Context, tokenURL string) (oauth2.TokenSource, error) {
	oauthConf := &clientcredentials.Config{
		ClientID:     c.ClientId,
		ClientSecret: c.ClientSecret,
		Scopes:       c.Scopes,
		TokenURL:     tokenURL,
	}

	return oauthConf.TokenSource(ctx), nil
}
