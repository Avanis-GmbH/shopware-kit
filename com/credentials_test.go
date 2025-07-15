package com

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewPasswordCredentials(t *testing.T) {
	username := "admin"
	password := "password123"
	scopes := []string{"write", "read"}

	creds := NewPasswordCredentials(username, password, scopes)

	if creds.Username != username {
		t.Errorf("NewPasswordCredentials() Username = %v, want %v", creds.Username, username)
	}

	if creds.Password != password {
		t.Errorf("NewPasswordCredentials() Password = %v, want %v", creds.Password, password)
	}

	if len(creds.Scopes) != len(scopes) {
		t.Errorf("NewPasswordCredentials() Scopes length = %v, want %v", len(creds.Scopes), len(scopes))
	}

	for i, scope := range scopes {
		if creds.Scopes[i] != scope {
			t.Errorf("NewPasswordCredentials() Scopes[%d] = %v, want %v", i, creds.Scopes[i], scope)
		}
	}
}

func TestPasswordCredentials_GetTokenSource(t *testing.T) {
	tests := []struct {
		name     string
		username string
		password string
		scopes   []string
		tokenURL string
		wantErr  bool
	}{
		{
			name:     "valid credentials",
			username: "admin",
			password: "password",
			scopes:   []string{"write"},
			tokenURL: "https://example.com/api/oauth/token",
			wantErr:  true, // Will fail without actual server
		},
		{
			name:     "empty username",
			username: "",
			password: "password",
			scopes:   []string{"write"},
			tokenURL: "https://example.com/api/oauth/token",
			wantErr:  true,
		},
		{
			name:     "invalid token URL",
			username: "admin",
			password: "password",
			scopes:   []string{"write"},
			tokenURL: "://invalid-url",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			creds := NewPasswordCredentials(tt.username, tt.password, tt.scopes)
			_, err := creds.GetTokenSource(context.Background(), tt.tokenURL)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetTokenSource() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewIntegrationCredentials(t *testing.T) {
	clientID := "test-client-id"
	clientSecret := "test-client-secret"
	scopes := []string{"write", "read"}

	creds := NewIntegrationCredentials(clientID, clientSecret, scopes)

	if creds.ClientID != clientID {
		t.Errorf("NewIntegrationCredentials() ClientID = %v, want %v", creds.ClientID, clientID)
	}

	if creds.ClientSecret != clientSecret {
		t.Errorf("NewIntegrationCredentials() ClientSecret = %v, want %v", creds.ClientSecret, clientSecret)
	}

	if len(creds.Scopes) != len(scopes) {
		t.Errorf("NewIntegrationCredentials() Scopes length = %v, want %v", len(creds.Scopes), len(scopes))
	}

	for i, scope := range scopes {
		if creds.Scopes[i] != scope {
			t.Errorf("NewIntegrationCredentials() Scopes[%d] = %v, want %v", i, creds.Scopes[i], scope)
		}
	}
}

func TestIntegrationCredentials_GetTokenSource(t *testing.T) {
	tests := []struct {
		name         string
		clientID     string
		clientSecret string
		scopes       []string
		tokenURL     string
		wantErr      bool
	}{
		{
			name:         "valid credentials",
			clientID:     "test-id",
			clientSecret: "test-secret",
			scopes:       []string{"write"},
			tokenURL:     "https://example.com/api/oauth/token",
			wantErr:      false,
		},
		{
			name:         "empty client ID",
			clientID:     "",
			clientSecret: "test-secret",
			scopes:       []string{"write"},
			tokenURL:     "https://example.com/api/oauth/token",
			wantErr:      false, // clientcredentials doesn't validate this immediately
		},
		{
			name:         "empty token URL",
			clientID:     "test-id",
			clientSecret: "test-secret",
			scopes:       []string{"write"},
			tokenURL:     "",
			wantErr:      false, // TokenSource is created but may fail on use
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			creds := NewIntegrationCredentials(tt.clientID, tt.clientSecret, tt.scopes)
			tokenSource, err := creds.GetTokenSource(context.Background(), tt.tokenURL)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetTokenSource() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && tokenSource == nil {
				t.Error("GetTokenSource() returned nil token source")
			}
		})
	}
}

func TestIntegrationCredentials_GetTokenSource_Error(t *testing.T) {
	tests := []struct {
		name         string
		clientID     string
		clientSecret string
		scopes       []string
		tokenURL     string
		wantErr      bool
	}{
		{
			name:         "valid parameters",
			clientID:     "test-id",
			clientSecret: "test-secret",
			scopes:       []string{"write"},
			tokenURL:     "https://example.com/api/oauth/token",
			wantErr:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			creds := NewIntegrationCredentials(tt.clientID, tt.clientSecret, tt.scopes)
			_, err := creds.GetTokenSource(context.Background(), tt.tokenURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTokenSource() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPasswordCredentials_GetTokenSource_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"access_token": "test-token",
			"token_type":   "Bearer",
			"expires_in":   3600,
		})
	}))
	defer server.Close()

	creds := NewPasswordCredentials("user", "pass", []string{"write"})
	tokenSource, err := creds.GetTokenSource(context.Background(), server.URL+"/token")
	if err != nil {
		t.Errorf("GetTokenSource() error = %v", err)
	}
	if tokenSource == nil {
		t.Error("GetTokenSource() returned nil token source")
	}
}

func TestIntegrationCredentials_GetTokenSource_Success(t *testing.T) {
	creds := NewIntegrationCredentials("client-id", "client-secret", []string{"write"})
	tokenSource, err := creds.GetTokenSource(context.Background(), "https://example.com/token")
	if err != nil {
		t.Errorf("GetTokenSource() error = %v", err)
	}
	if tokenSource == nil {
		t.Error("GetTokenSource() returned nil token source")
	}
}

func TestOAuthCredentials_Interface(t *testing.T) {
	// Test that both credential types implement the OAuthCredentials interface
	var _ OAuthCredentials = NewPasswordCredentials("user", "pass", []string{"write"})
	var _ OAuthCredentials = NewIntegrationCredentials("id", "secret", []string{"write"})
}
