package com

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

// mockHTTPClient creates a mock HTTP client with a predefined response
func mockHTTPClient(statusCode int, body string, headers map[string]string) *http.Client {
	return &http.Client{
		Transport: &mockTransport{
			statusCode: statusCode,
			body:       body,
			headers:    headers,
		},
	}
}

type mockTransport struct {
	statusCode int
	body       string
	headers    map[string]string
	requests   []*http.Request
}

func (m *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	m.requests = append(m.requests, req)

	resp := &http.Response{
		StatusCode: m.statusCode,
		Body:       io.NopCloser(strings.NewReader(m.body)),
		Header:     make(http.Header),
		Request:    req,
	}

	for k, v := range m.headers {
		resp.Header.Set(k, v)
	}

	return resp, nil
}

func TestNewClient(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/oauth/token" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"access_token": "test-token",
				"token_type":   "Bearer",
				"expires_in":   3600,
			})
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	tests := []struct {
		name        string
		shopURL     string
		credentials OAuthCredentials
		httpClient  *http.Client
		wantErr     bool
	}{
		{
			name:        "valid integration credentials",
			shopURL:     server.URL,
			credentials: NewIntegrationCredentials("test-id", "test-secret", []string{"write"}),
			httpClient:  nil,
			wantErr:     false,
		},
		{
			name:        "valid password credentials",
			shopURL:     server.URL,
			credentials: NewPasswordCredentials("admin", "password", []string{"write"}),
			httpClient:  nil,
			wantErr:     false,
		},
		{
			name:        "invalid URL",
			shopURL:     "://invalid-url",
			credentials: NewIntegrationCredentials("test-id", "test-secret", []string{"write"}),
			httpClient:  nil,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewClient(context.Background(), tt.shopURL, tt.credentials, tt.httpClient)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && client == nil {
				t.Errorf("NewClient() returned nil client")
			}
			if !tt.wantErr && client != nil {
				if client.remote != tt.shopURL {
					t.Errorf("NewClient() remote = %v, want %v", client.remote, tt.shopURL)
				}
			}
		})
	}
}

func TestClient_NewRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"access_token": "test-token",
			"token_type":   "Bearer",
			"expires_in":   3600,
		})
	}))
	defer server.Close()

	client, err := NewClient(context.Background(), server.URL, NewIntegrationCredentials("test", "test", []string{"write"}), nil)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := NewAPIContext(context.Background())

	tests := []struct {
		name    string
		method  string
		path    string
		options map[string]string
		body    interface{}
		wantErr bool
	}{
		{
			name:    "GET request",
			method:  "GET",
			path:    "/api/product",
			options: nil,
			body:    nil,
			wantErr: false,
		},
		{
			name:    "POST request with body",
			method:  "POST",
			path:    "/api/product",
			options: nil,
			body:    map[string]string{"name": "test"},
			wantErr: false,
		},
		{
			name:    "GET request with query params",
			method:  "GET",
			path:    "/api/product",
			options: map[string]string{"limit": "10", "page": "1"},
			body:    nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := client.NewRequest(ctx, tt.method, tt.path, tt.options, tt.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if req.Method != tt.method {
					t.Errorf("NewRequest() method = %v, want %v", req.Method, tt.method)
				}

				expectedURL, _ := url.JoinPath(server.URL, tt.path)
				if !strings.HasPrefix(req.URL.String(), expectedURL) {
					t.Errorf("NewRequest() URL = %v, want to start with %v", req.URL.String(), expectedURL)
				}

				if req.Header.Get("Accept") != "application/json" {
					t.Errorf("NewRequest() Accept header = %v, want application/json", req.Header.Get("Accept"))
				}

				if req.Header.Get("sw-language-id") != ctx.LanguageID {
					t.Errorf("NewRequest() sw-language-id header = %v, want %v", req.Header.Get("sw-language-id"), ctx.LanguageID)
				}

				if tt.body != nil && req.Header.Get("Content-Type") != "application/json" {
					t.Errorf("NewRequest() Content-Type header = %v, want application/json", req.Header.Get("Content-Type"))
				}

				if tt.options != nil {
					for key, value := range tt.options {
						if req.URL.Query().Get(key) != value {
							t.Errorf("NewRequest() query param %s = %v, want %v", key, req.URL.Query().Get(key), value)
						}
					}
				}
			}
		})
	}
}

func TestClient_NewRawRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"access_token": "test-token",
			"token_type":   "Bearer",
			"expires_in":   3600,
		})
	}))
	defer server.Close()

	client, err := NewClient(context.Background(), server.URL, NewIntegrationCredentials("test", "test", []string{"write"}), nil)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := NewAPIContext(context.Background())
	ctx.SkipFlows = true

	body := strings.NewReader("test body")
	req, err := client.NewRawRequest(ctx, "POST", "/api/test", map[string]string{"param": "value"}, body)

	if err != nil {
		t.Fatalf("NewRawRequest() error = %v", err)
	}

	if req.Header.Get("sw-skip-trigger-flow") != "1" {
		t.Errorf("NewRawRequest() skip flows header not set correctly")
	}

	if req.Header.Get("User-Agent") != userAgent {
		t.Errorf("NewRawRequest() User-Agent = %v, want %v", req.Header.Get("User-Agent"), userAgent)
	}
}

func TestClient_Do(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		responseBody   string
		responseStruct interface{}
		wantErr        bool
	}{
		{
			name:           "successful JSON response",
			statusCode:     200,
			responseBody:   `{"name": "test", "id": "123"}`,
			responseStruct: &map[string]interface{}{},
			wantErr:        false,
		},
		{
			name:           "empty response",
			statusCode:     204,
			responseBody:   "",
			responseStruct: nil,
			wantErr:        false,
		},
		{
			name:           "invalid JSON response",
			statusCode:     200,
			responseBody:   `{"invalid": json}`,
			responseStruct: &map[string]interface{}{},
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path == "/api/oauth/token" {
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(map[string]interface{}{
						"access_token": "test-token",
						"token_type":   "Bearer",
						"expires_in":   3600,
					})
					return
				}
				w.WriteHeader(tt.statusCode)
				w.Write([]byte(tt.responseBody))
			}))
			defer server.Close()

			client, err := NewClient(context.Background(), server.URL, NewIntegrationCredentials("test", "test", []string{"write"}), nil)
			if err != nil {
				t.Fatalf("Failed to create client: %v", err)
			}

			ctx := NewAPIContext(context.Background())
			req, err := client.NewRequest(ctx, "GET", "/api/test", nil, nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			_, err = client.Do(ctx.Context, req, tt.responseStruct)
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_BareDo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/oauth/token" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"access_token": "test-token",
				"token_type":   "Bearer",
				"expires_in":   3600,
			})
			return
		}
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	}))
	defer server.Close()

	client, err := NewClient(context.Background(), server.URL, NewIntegrationCredentials("test", "test", []string{"write"}), nil)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := NewAPIContext(context.Background())
	req, err := client.NewRequest(ctx, "GET", "/api/test", nil, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	resp, err := client.BareDo(ctx.Context, req)
	if err != nil {
		t.Errorf("BareDo() error = %v", err)
		return
	}
	if resp.StatusCode != 200 {
		t.Errorf("BareDo() status = %v, want 200", resp.StatusCode)
	}
}

func TestClient_BareDo_ContextCancelled(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/oauth/token" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"access_token": "test-token",
				"token_type":   "Bearer",
				"expires_in":   3600,
			})
			return
		}
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	}))
	defer server.Close()

	client, err := NewClient(context.Background(), server.URL, NewIntegrationCredentials("test", "test", []string{"write"}), nil)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	apiCtx := NewAPIContext(ctx)
	req, err := client.NewRequest(apiCtx, "GET", "/api/test", nil, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	_, err = client.BareDo(ctx, req)
	if err == nil {
		t.Error("BareDo() should return error for cancelled context")
	}
}

func TestClient_BareDo_NilContext(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/oauth/token" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"access_token": "test-token",
				"token_type":   "Bearer",
				"expires_in":   3600,
			})
			return
		}
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	}))
	defer server.Close()

	client, err := NewClient(context.Background(), server.URL, NewIntegrationCredentials("test", "test", []string{"write"}), nil)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := NewAPIContext(context.Background())
	req, err := client.NewRequest(ctx, "GET", "/api/test", nil, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	_, err = client.BareDo(nil, req)
	if err == nil {
		t.Error("BareDo() should return error for nil context")
	}
}

func TestClient_DoWithWriter(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/oauth/token" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"access_token": "test-token",
				"token_type":   "Bearer",
				"expires_in":   3600,
			})
			return
		}
		w.Write([]byte("test data"))
	}))
	defer server.Close()

	client, err := NewClient(context.Background(), server.URL, NewIntegrationCredentials("test", "test", []string{"write"}), nil)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := NewAPIContext(context.Background())
	req, err := client.NewRequest(ctx, "GET", "/api/test", nil, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	var buf bytes.Buffer
	_, err = client.Do(ctx.Context, req, &buf)
	if err != nil {
		t.Errorf("Do() with writer error = %v", err)
	}

	if buf.String() != "test data" {
		t.Errorf("Do() with writer got = %v, want 'test data'", buf.String())
	}
}

func TestClient_WithCustomResponseHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/oauth/token" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"access_token": "test-token",
				"token_type":   "Bearer",
				"expires_in":   3600,
			})
			return
		}
		w.WriteHeader(400)
		w.Write([]byte("custom error"))
	}))
	defer server.Close()

	client, err := NewClient(context.Background(), server.URL, NewIntegrationCredentials("test", "test", []string{"write"}), nil)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	customErrorHandled := false
	client.ResponseHandler = func(resp *http.Response) error {
		customErrorHandled = true
		return errors.New("custom error handler")
	}

	ctx := NewAPIContext(context.Background())
	req, err := client.NewRequest(ctx, "GET", "/api/test", nil, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	_, err = client.Do(ctx.Context, req, nil)
	if err == nil {
		t.Error("Do() should return error when ResponseHandler is set")
	}

	if !customErrorHandled {
		t.Error("Custom response handler was not called")
	}
}

func TestClient_checkResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"access_token": "test-token",
			"token_type":   "Bearer",
			"expires_in":   3600,
		})
	}))
	defer server.Close()

	client, err := NewClient(context.Background(), server.URL, NewIntegrationCredentials("test", "test", []string{"write"}), nil)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	tests := []struct {
		name       string
		statusCode int
		body       string
		wantErr    bool
	}{
		{
			name:       "success response",
			statusCode: 200,
			body:       `{"data": "success"}`,
			wantErr:    false,
		},
		{
			name:       "created response",
			statusCode: 201,
			body:       `{"data": "created"}`,
			wantErr:    false,
		},
		{
			name:       "error response",
			statusCode: 400,
			body:       `{"errors": [{"status": "400", "code": "INVALID_REQUEST", "title": "Bad Request", "detail": "Invalid data"}]}`,
			wantErr:    true,
		},
		{
			name:       "server error",
			statusCode: 500,
			body:       `{"errors": [{"status": "500", "code": "INTERNAL_ERROR", "title": "Internal Server Error", "detail": "Something went wrong"}]}`,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				StatusCode: tt.statusCode,
				Body:       io.NopCloser(strings.NewReader(tt.body)),
			}

			err := client.checkResponse(resp)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkResponse() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr && err != nil {
				if apiErr, ok := err.(*APIError); ok {
					if apiErr.StatusCode != tt.statusCode {
						t.Errorf("checkResponse() status code = %v, want %v", apiErr.StatusCode, tt.statusCode)
					}
				}
			}
		})
	}
}

func TestSetUserAgent(t *testing.T) {
	originalUserAgent := userAgent
	defer func() {
		userAgent = originalUserAgent
	}()

	testUserAgent := "test-user-agent"
	SetUserAgent(testUserAgent)

	if userAgent != testUserAgent {
		t.Errorf("SetUserAgent() userAgent = %v, want %v", userAgent, testUserAgent)
	}
}

func TestNewAPIContext_SkipFlows(t *testing.T) {
	ctx := context.Background()
	apiCtx := NewAPIContext(ctx)

	if apiCtx.SkipFlows {
		t.Errorf("NewAPIContext() SkipFlows = %v, want false", apiCtx.SkipFlows)
	}

	apiCtx.SkipFlows = true
	if !apiCtx.SkipFlows {
		t.Errorf("APIContext.SkipFlows = %v, want true", apiCtx.SkipFlows)
	}
}

func TestClient_authorize_InvalidURL(t *testing.T) {
	client := &Client{
		remote:      "://invalid-url", // Invalid URL scheme
		ctx:         context.Background(),
		credentials: NewIntegrationCredentials("test", "test", []string{"write"}),
	}

	err := client.authorize()
	if err == nil {
		t.Error("authorize() should return error for invalid URL")
	}
}

func TestClient_authorize_CredentialsError(t *testing.T) {
	mockCreds := &mockFailingCredentials{}

	client := &Client{
		remote:      "https://example.com",
		ctx:         context.Background(),
		credentials: mockCreds,
	}

	err := client.authorize()
	if err == nil {
		t.Error("authorize() should return error when credentials.GetTokenSource fails")
	}
}

type mockFailingCredentials struct{}

func (m *mockFailingCredentials) GetTokenSource(ctx context.Context, tokenURL string) (oauth2.TokenSource, error) {
	return nil, errors.New("mock credentials error")
}

func TestClient_NewRequest_JSONEncodeError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"access_token": "test-token",
			"token_type":   "Bearer",
			"expires_in":   3600,
		})
	}))
	defer server.Close()

	client, err := NewClient(context.Background(), server.URL, NewIntegrationCredentials("test", "test", []string{"write"}), nil)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := NewAPIContext(context.Background())

	invalidBody := make(chan int)

	_, err = client.NewRequest(ctx, "POST", "/api/test", nil, invalidBody)
	if err == nil {
		t.Error("NewRequest() should return error when JSON encoding fails")
	}
}

func TestClient_NewRawRequest_InvalidURL(t *testing.T) {
	client := &Client{
		remote: "://invalid-url",
	}

	ctx := NewAPIContext(context.Background())
	_, err := client.NewRawRequest(ctx, "GET", "/api/test", nil, nil)
	if err == nil {
		t.Error("NewRawRequest() should return error for invalid URL")
	}
}

func TestClient_NewRawRequest_InvalidPath(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"access_token": "test-token",
			"token_type":   "Bearer",
			"expires_in":   3600,
		})
	}))
	defer server.Close()

	client, err := NewClient(context.Background(), server.URL, NewIntegrationCredentials("test", "test", []string{"write"}), nil)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := NewAPIContext(context.Background())

	_, err = client.NewRawRequest(ctx, "INVALID\nMETHOD", "/api/test", nil, nil)
	if err == nil {
		t.Error("NewRawRequest() should return error for invalid HTTP method")
	}
}

func TestClient_checkResponse_ReadBodyError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/oauth/token" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"access_token": "test-token",
				"token_type":   "Bearer",
				"expires_in":   3600,
			})
			return
		}
		w.WriteHeader(400)
		w.Header().Set("Content-Length", "1000") // Claim 1000 bytes
		w.Write([]byte("short"))                 // But only send 5 bytes
	}))
	defer server.Close()

	client, err := NewClient(context.Background(), server.URL, NewIntegrationCredentials("test", "test", []string{"write"}), nil)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := NewAPIContext(context.Background())
	req, err := client.NewRequest(ctx, "GET", "/api/test", nil, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// This should trigger an error in checkResponse's io.ReadAll call
	_, err = client.BareDo(ctx.Context, req)
	if err == nil {
		t.Error("BareDo() should return error when response body reading fails")
	}
}

func TestClient_BareDo_HTTPClientError(t *testing.T) {
	// Create a client with a mock HTTP client that always fails
	client := &Client{
		remote: "https://example.com",
		client: &http.Client{
			Transport: &mockFailingTransport{},
		},
	}

	ctx := context.Background()
	req, _ := http.NewRequestWithContext(ctx, "GET", "https://example.com/test", nil)

	_, err := client.BareDo(ctx, req)
	if err == nil {
		t.Error("BareDo() should return error when HTTP client fails")
	}
}

// Mock transport that always fails
type mockFailingTransport struct{}

func (m *mockFailingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return nil, errors.New("mock HTTP client error")
}

func TestClient_authorize_WithExistingClient(t *testing.T) {
	// Test authorize function when client already has an HTTP client set
	existingClient := &http.Client{}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"access_token": "test-token",
			"token_type":   "Bearer",
			"expires_in":   3600,
		})
	}))
	defer server.Close()

	client := &Client{
		remote:      server.URL,
		ctx:         context.Background(),
		credentials: NewIntegrationCredentials("test", "test", []string{"write"}),
		client:      existingClient, // Pre-existing client
	}

	err := client.authorize()
	if err != nil {
		t.Errorf("authorize() error = %v", err)
	}

	// Verify that the client was replaced with an OAuth2 client
	if client.client == existingClient {
		t.Error("authorize() should replace the existing client with OAuth2 client")
	}
}
