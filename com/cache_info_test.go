package com

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Clear(t *testing.T) {
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

		if r.URL.Path == "/api/_action/cache" && r.Method == "DELETE" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client, err := NewClient(context.Background(), server.URL, NewIntegrationCredentials("test", "test", []string{"write"}), nil)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := NewAPIContext(context.Background())

	resp, err := client.Clear(ctx)
	if err != nil {
		t.Errorf("Clear() error = %v", err)
		return
	}

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Clear() status = %v, want %v", resp.StatusCode, http.StatusNoContent)
	}
}

func TestClient_Clear_Error(t *testing.T) {
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

		if r.URL.Path == "/api/_action/cache" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"errors": []map[string]interface{}{
					{
						"status": "403",
						"code":   "FORBIDDEN",
						"detail": "Cache clear not allowed",
					},
				},
			})
			return
		}

		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client, err := NewClient(context.Background(), server.URL, NewIntegrationCredentials("test", "test", []string{"write"}), nil)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := NewAPIContext(context.Background())

	_, err = client.Clear(ctx)
	if err == nil {
		t.Error("Clear() should return error for forbidden response")
	}
}

func TestClient_Info(t *testing.T) {
	expectedInfo := InfoResponse{
		Version:         "6.4.0.0",
		VersionRevision: "abc123",
		AdminWorker: struct {
			EnableAdminWorker bool     `json:"enableAdminWorker"`
			Transports        []string `json:"transports"`
		}{
			EnableAdminWorker: true,
			Transports:        []string{"async"},
		},
		Bundles: map[string]infoResponseBundle{
			"TestBundle": {
				CSS: []string{"test.css"},
				Js:  []string{"test.js"},
			},
		},
		Settings: struct {
			EnableURLFeature bool `json:"enableUrlFeature"`
		}{
			EnableURLFeature: true,
		},
	}

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

		if r.URL.Path == "/api/_info/config" && r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(expectedInfo)
			return
		}

		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client, err := NewClient(context.Background(), server.URL, NewIntegrationCredentials("test", "test", []string{"write"}), nil)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := NewAPIContext(context.Background())

	info, resp, err := client.Info(ctx)
	if err != nil {
		t.Errorf("Info() error = %v", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Info() status = %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if info == nil {
		t.Fatal("Info() returned nil info")
	}

	if info.Version != expectedInfo.Version {
		t.Errorf("Info() Version = %v, want %v", info.Version, expectedInfo.Version)
	}

	if info.VersionRevision != expectedInfo.VersionRevision {
		t.Errorf("Info() VersionRevision = %v, want %v", info.VersionRevision, expectedInfo.VersionRevision)
	}

	if info.AdminWorker.EnableAdminWorker != expectedInfo.AdminWorker.EnableAdminWorker {
		t.Errorf("Info() AdminWorker.EnableAdminWorker = %v, want %v",
			info.AdminWorker.EnableAdminWorker, expectedInfo.AdminWorker.EnableAdminWorker)
	}

	if len(info.AdminWorker.Transports) != len(expectedInfo.AdminWorker.Transports) {
		t.Errorf("Info() AdminWorker.Transports length = %v, want %v",
			len(info.AdminWorker.Transports), len(expectedInfo.AdminWorker.Transports))
	}

	if info.Settings.EnableURLFeature != expectedInfo.Settings.EnableURLFeature {
		t.Errorf("Info() Settings.EnableURLFeature = %v, want %v",
			info.Settings.EnableURLFeature, expectedInfo.Settings.EnableURLFeature)
	}
}

func TestClient_Info_ErrorResponse(t *testing.T) {
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

		if r.URL.Path == "/api/_info/config" && r.Method == "GET" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"errors": []map[string]interface{}{
					{
						"status": "403",
						"code":   "FORBIDDEN",
						"detail": "Access to info endpoint denied",
					},
				},
			})
			return
		}

		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client, err := NewClient(context.Background(), server.URL, NewIntegrationCredentials("test", "test", []string{"write"}), nil)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := NewAPIContext(context.Background())
	_, _, err = client.Info(ctx)
	if err == nil {
		t.Error("Info() should return error when API returns error")
	}
}

func TestIsCloudShop(t *testing.T) {
	tests := []struct {
		name     string
		info     InfoResponse
		expected bool
	}{
		{
			name: "cloud shop with SaasRufus bundle",
			info: InfoResponse{
				Version: "6.4.15.0",
				Bundles: map[string]infoResponseBundle{
					"SaasRufus": {
						CSS: []string{"test.css"},
						Js:  []string{"test.js"},
					},
				},
				Settings: struct {
					EnableURLFeature bool `json:"enableUrlFeature"`
				}{
					EnableURLFeature: true,
				},
			},
			expected: true,
		},
		{
			name: "non-cloud shop without SaasRufus bundle",
			info: InfoResponse{
				Version: "6.4.15.0-dev",
				Bundles: map[string]infoResponseBundle{
					"OtherBundle": {
						CSS: []string{"other.css"},
						Js:  []string{"other.js"},
					},
				},
				Settings: struct {
					EnableURLFeature bool `json:"enableUrlFeature"`
				}{
					EnableURLFeature: false,
				},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.info.IsCloudShop()
			if result != tt.expected {
				t.Errorf("IsCloudShop() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestInfoResponse_JSONSerialization(t *testing.T) {
	info := InfoResponse{
		Version:         "6.4.0.0",
		VersionRevision: "abc123",
		AdminWorker: struct {
			EnableAdminWorker bool     `json:"enableAdminWorker"`
			Transports        []string `json:"transports"`
		}{
			EnableAdminWorker: true,
			Transports:        []string{"async", "sync"},
		},
		Bundles: map[string]infoResponseBundle{
			"TestBundle": {
				CSS: []string{"test.css", "another.css"},
				Js:  []string{"test.js"},
			},
		},
		Settings: struct {
			EnableURLFeature bool `json:"enableUrlFeature"`
		}{
			EnableURLFeature: false,
		},
	}

	data, err := json.Marshal(info)
	if err != nil {
		t.Fatalf("Failed to marshal InfoResponse: %v", err)
	}

	var unmarshaled InfoResponse
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal InfoResponse: %v", err)
	}

	if info.Version != unmarshaled.Version {
		t.Errorf("Version not preserved: got %v, want %v", unmarshaled.Version, info.Version)
	}

	if len(info.AdminWorker.Transports) != len(unmarshaled.AdminWorker.Transports) {
		t.Errorf("AdminWorker.Transports length not preserved: got %v, want %v",
			len(unmarshaled.AdminWorker.Transports), len(info.AdminWorker.Transports))
	}

	if len(info.Bundles) != len(unmarshaled.Bundles) {
		t.Errorf("Bundles length not preserved: got %v, want %v",
			len(unmarshaled.Bundles), len(info.Bundles))
	}

	testBundle, exists := unmarshaled.Bundles["TestBundle"]
	if !exists {
		t.Error("TestBundle not found in unmarshaled bundles")
	} else {
		if len(testBundle.CSS) != 2 {
			t.Errorf("TestBundle CSS length not preserved: got %v, want 2", len(testBundle.CSS))
		}
		if len(testBundle.Js) != 1 {
			t.Errorf("TestBundle Js length not preserved: got %v, want 1", len(testBundle.Js))
		}
	}
}

func TestInfoResponseBundle_Structure(t *testing.T) {
	bundle := infoResponseBundle{
		CSS: []string{"style1.css", "style2.css"},
		Js:  []string{"script1.js", "script2.js", "script3.js"},
	}

	data, err := json.Marshal(bundle)
	if err != nil {
		t.Fatalf("Failed to marshal infoResponseBundle: %v", err)
	}

	var unmarshaled infoResponseBundle
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal infoResponseBundle: %v", err)
	}

	if len(bundle.CSS) != len(unmarshaled.CSS) {
		t.Errorf("CSS length not preserved: got %v, want %v", len(unmarshaled.CSS), len(bundle.CSS))
	}

	if len(bundle.Js) != len(unmarshaled.Js) {
		t.Errorf("Js length not preserved: got %v, want %v", len(unmarshaled.Js), len(bundle.Js))
	}

	for i, css := range bundle.CSS {
		if unmarshaled.CSS[i] != css {
			t.Errorf("CSS[%d] not preserved: got %v, want %v", i, unmarshaled.CSS[i], css)
		}
	}

	for i, js := range bundle.Js {
		if unmarshaled.Js[i] != js {
			t.Errorf("Js[%d] not preserved: got %v, want %v", i, unmarshaled.Js[i], js)
		}
	}
}

func TestClient_Clear_InvalidResponse(t *testing.T) {
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

		if r.URL.Path == "/api/_action/cache" && r.Method == "DELETE" {
			w.WriteHeader(http.StatusOK) // Wrong status code for cache clear
			w.Write([]byte("OK"))
			return
		}

		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client, err := NewClient(context.Background(), server.URL, NewIntegrationCredentials("test", "test", []string{"write"}), nil)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := NewAPIContext(context.Background())
	resp, err := client.Clear(ctx)
	if err != nil {
		t.Errorf("Clear() error = %v", err)
	}

	// Should still return the response even if status code is not 204
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Clear() status = %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestClient_Clear_RequestCreationError(t *testing.T) {
	// Create a client that will fail on NewRequest due to invalid URL
	client := &Client{
		remote: "://invalid-url", // This will cause url.JoinPath to fail
	}

	ctx := NewAPIContext(context.Background())
	_, err := client.Clear(ctx)
	if err == nil {
		t.Error("Clear() should return error when NewRequest fails")
	}
}

func TestClient_Info_RequestCreationError(t *testing.T) {
	// Create a client that will fail on NewRequest due to invalid URL
	client := &Client{
		remote: "://invalid-url", // This will cause url.JoinPath to fail
	}

	ctx := NewAPIContext(context.Background())
	_, _, err := client.Info(ctx)
	if err == nil {
		t.Error("Info() should return error when NewRequest fails")
	}
}
