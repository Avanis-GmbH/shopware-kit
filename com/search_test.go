package com

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestClient_Search(t *testing.T) {
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

		if r.URL.Path == "/api/search/mock-entity" && r.Method == "POST" {
			response := MockEntityCollection{
				EntityCollection: EntityCollection{
					Total: 2,
				},
				Data: []MockEntity{
					{ID: "1", Name: "Entity 1"},
					{ID: "2", Name: "Entity 2"},
				},
			}
			json.NewEncoder(w).Encode(response)
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
	criteria := Criteria{
		Limit: 10,
		Page:  1,
	}

	var collection MockEntityCollection
	err = client.Search(ctx, criteria, &collection)
	if err != nil {
		t.Errorf("Search() error = %v", err)
	}

	if collection.Total != 2 {
		t.Errorf("Search() total = %v, want 2", collection.Total)
	}

	if len(collection.Data) != 2 {
		t.Errorf("Search() data length = %v, want 2", len(collection.Data))
		return
	}

	if collection.Data[0].Name != "Entity 1" {
		t.Errorf("Search() first entity name = %v, want 'Entity 1'", collection.Data[0].Name)
	}
}

func TestClient_SearchIDs(t *testing.T) {
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

		if r.URL.Path == "/api/search-ids/mock-entity" && r.Method == "POST" {
			response := SearchIDsResponse{
				Total: 3,
				Data:  []string{"id1", "id2", "id3"},
			}
			json.NewEncoder(w).Encode(response)
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
	criteria := Criteria{
		Limit: 100,
	}

	var mockEntity MockEntity
	response, err := client.SearchIDs(ctx, criteria, mockEntity)
	if err != nil {
		t.Errorf("SearchIDs() error = %v", err)
	}

	if response.Total != 3 {
		t.Errorf("SearchIDs() total = %v, want 3", response.Total)
	}

	if len(response.Data) != 3 {
		t.Errorf("SearchIDs() data length = %v, want 3", len(response.Data))
	}

	expectedIDs := []string{"id1", "id2", "id3"}
	if !reflect.DeepEqual(response.Data, expectedIDs) {
		t.Errorf("SearchIDs() data = %v, want %v", response.Data, expectedIDs)
	}
}

func TestSearchIDsResponse_JSONSerialization(t *testing.T) {
	response := SearchIDsResponse{
		Total: 5,
		Data:  []string{"id1", "id2", "id3", "id4", "id5"},
	}

	data, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("Failed to marshal SearchIDsResponse: %v", err)
	}

	var unmarshaled SearchIDsResponse
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal SearchIDsResponse: %v", err)
	}

	if response.Total != unmarshaled.Total {
		t.Errorf("Total not preserved: got %v, want %v", unmarshaled.Total, response.Total)
	}

	if !reflect.DeepEqual(response.Data, unmarshaled.Data) {
		t.Errorf("Data not preserved: got %v, want %v", unmarshaled.Data, response.Data)
	}
}

func TestClient_SearchAll(t *testing.T) {
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

		if r.URL.Path == "/api/search/mock-entity" && r.Method == "POST" {
			var criteria Criteria
			json.NewDecoder(r.Body).Decode(&criteria)

			if criteria.Page == 2 {
				response := map[string]interface{}{
					"total": 3,
					"data": []interface{}{
						map[string]interface{}{"id": "1", "name": "Entity 1"},
						map[string]interface{}{"id": "2", "name": "Entity 2"},
					},
				}
				json.NewEncoder(w).Encode(response)
				return
			} else if criteria.Page == 3 {
				response := map[string]interface{}{
					"total": 3,
					"data": []interface{}{
						map[string]interface{}{"id": "3", "name": "Entity 3"},
					},
				}
				json.NewEncoder(w).Encode(response)
				return
			} else {
				response := map[string]interface{}{
					"total": 3,
					"data":  []interface{}{},
				}
				json.NewEncoder(w).Encode(response)
				return
			}
		}

		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client, err := NewClient(context.Background(), server.URL, NewIntegrationCredentials("test", "test", []string{"write"}), nil)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := NewAPIContext(context.Background())
	criteria := Criteria{
		Limit: 2,
		Page:  1,
	}

	var collection MockEntityCollection
	err = client.SearchAll(ctx, criteria, &collection)
	if err != nil {
		t.Errorf("SearchAll() error = %v", err)
	}

	if collection.Total != 3 {
		t.Errorf("SearchAll() total = %v, want 3", collection.Total)
	}

	if len(collection.getData()) != 3 {
		t.Errorf("SearchAll() data length = %v, want 3", len(collection.getData()))
	}
}

func TestClient_SearchAll_DefaultLimitAndPage(t *testing.T) {
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

		if r.URL.Path == "/api/search/mock-entity" && r.Method == "POST" {
			var criteria Criteria
			json.NewDecoder(r.Body).Decode(&criteria)

			if criteria.Limit != 50 {
				t.Errorf("SearchAll() default limit = %v, want 50", criteria.Limit)
			}

			response := MockEntityCollection{
				EntityCollection: EntityCollection{
					Total: 0,
				},
				Data: []MockEntity{},
			}
			json.NewEncoder(w).Encode(response)
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
	criteria := Criteria{} // Empty criteria to test defaults

	var collection MockEntityCollection
	err = client.SearchAll(ctx, criteria, &collection)
	if err != nil {
		t.Errorf("SearchAll() error = %v", err)
	}
}

func TestClient_SearchAll_Error(t *testing.T) {
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

		if r.URL.Path == "/api/search/mock-entity" && r.Method == "POST" {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"errors": []map[string]interface{}{
					{
						"status": "500",
						"code":   "INTERNAL_ERROR",
						"detail": "Search failed",
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
	criteria := Criteria{
		Limit: 10,
		Page:  1,
	}

	var collection MockEntityCollection
	err = client.SearchAll(ctx, criteria, &collection)
	if err == nil {
		t.Error("SearchAll() should return error when search fails")
	}
}

func TestClient_Search_Error(t *testing.T) {
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

		if r.URL.Path == "/api/search/mock-entity" && r.Method == "POST" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"errors": []map[string]interface{}{
					{
						"status": "400",
						"code":   "BAD_REQUEST",
						"detail": "Invalid criteria",
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
	criteria := Criteria{
		Limit: 10,
		Page:  1,
	}

	var collection MockEntityCollection
	err = client.Search(ctx, criteria, &collection)
	if err == nil {
		t.Error("Search() should return error when API returns error")
	}
}

func TestClient_SearchIDs_Error(t *testing.T) {
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

		if r.URL.Path == "/api/search-ids/mock-entity" && r.Method == "POST" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"errors": []map[string]interface{}{
					{
						"status": "403",
						"code":   "FORBIDDEN",
						"detail": "Access denied",
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
	criteria := Criteria{
		Limit: 100,
	}

	var mockEntity MockEntity
	_, err = client.SearchIDs(ctx, criteria, mockEntity)
	if err == nil {
		t.Error("SearchIDs() should return error when API returns error")
	}
}

func TestClient_Search_URLJoinError(t *testing.T) {
	client := &Client{
		remote: "://invalid-url",
	}

	ctx := NewAPIContext(context.Background())
	criteria := Criteria{Limit: 10}
	var collection MockEntityCollection

	err := client.Search(ctx, criteria, &collection)
	if err == nil {
		t.Error("Search() should return error when url.JoinPath fails")
	}
}

func TestClient_Search_NewRequestError(t *testing.T) {
	client := &Client{
		remote: "://invalid-url", // This will make url.JoinPath in NewRawRequest fail
	}

	ctx := NewAPIContext(context.Background())
	criteria := Criteria{Limit: 10}
	var collection MockEntityCollection

	err := client.Search(ctx, criteria, &collection)
	if err == nil {
		t.Error("Search() should return error when NewRequest fails")
	}
}

func TestClient_SearchIDs_URLJoinError(t *testing.T) {
	client := &Client{
		remote: "://invalid-url",
	}

	ctx := NewAPIContext(context.Background())
	criteria := Criteria{Limit: 10}
	var entity MockEntity

	_, err := client.SearchIDs(ctx, criteria, entity)
	if err == nil {
		t.Error("SearchIDs() should return error when url.JoinPath fails")
	}
}

func TestClient_SearchIDs_NewRequestError(t *testing.T) {
	client := &Client{
		remote: "://invalid-url", // This will make NewRawRequest fail
	}

	ctx := NewAPIContext(context.Background())
	criteria := Criteria{Limit: 10}
	var entity MockEntity

	_, err := client.SearchIDs(ctx, criteria, entity)
	if err == nil {
		t.Error("SearchIDs() should return error when NewRequest fails")
	}
}
