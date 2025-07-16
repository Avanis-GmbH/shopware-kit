package com

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestSyncOperation_JSONSerialization(t *testing.T) {
	operation := SyncOperation{
		Entity:  "product",
		Action:  "upsert",
		Payload: map[string]interface{}{"name": "Test Product", "id": "123"},
	}

	data, err := json.Marshal(operation)
	if err != nil {
		t.Fatalf("Failed to marshal SyncOperation: %v", err)
	}

	var unmarshaled SyncOperation
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal SyncOperation: %v", err)
	}

	if operation.Entity != unmarshaled.Entity {
		t.Errorf("Entity not preserved: got %v, want %v", unmarshaled.Entity, operation.Entity)
	}

	if operation.Action != unmarshaled.Action {
		t.Errorf("Action not preserved: got %v, want %v", unmarshaled.Action, operation.Action)
	}
}

func TestClient_Sync(t *testing.T) {
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

		if r.URL.Path == "/api/_action/sync" && r.Method == "POST" {
			var payload map[string]SyncOperation
			err := json.NewDecoder(r.Body).Decode(&payload)
			if err != nil {
				t.Errorf("Failed to decode sync payload: %v", err)
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": true,
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
	payload := map[string]SyncOperation{
		"product": {
			Entity:  "product",
			Action:  "upsert",
			Payload: []map[string]interface{}{{"name": "Test Product"}},
		},
	}

	resp, err := client.Sync(ctx, payload)
	if err != nil {
		t.Errorf("Sync() error = %v", err)
		return // Don't try to access resp if there was an error
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Sync() status = %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestClient_Upsert(t *testing.T) {
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

		if r.URL.Path == "/api/_action/sync" && r.Method == "POST" {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": true,
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

	entities := []MockEntity{
		{ID: "1", Name: "Entity 1"},
		{ID: "2", Name: "Entity 2"},
	}

	resp, err := client.Upsert(ctx, entities)
	if err != nil {
		t.Errorf("Upsert() error = %v", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Upsert() status = %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestClient_Upsert_Errors(t *testing.T) {
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

	entity := MockEntity{ID: "1", Name: "Single Entity"}
	_, err = client.Upsert(ctx, entity)
	if err == nil {
		t.Error("Upsert() should return error for non-slice entity")
	}

	emptySlice := []MockEntity{}
	_, err = client.Upsert(ctx, emptySlice)
	if err == nil {
		t.Error("Upsert() should return error for empty slice (unknown entity)")
	}
}

func TestClient_Delete(t *testing.T) {
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

		if r.URL.Path == "/api/_action/sync" && r.Method == "POST" {
			var payload map[string]SyncOperation
			err := json.NewDecoder(r.Body).Decode(&payload)
			if err != nil {
				t.Errorf("Failed to decode delete payload: %v", err)
			}

			for _, operation := range payload {
				if operation.Action != "delete" {
					t.Errorf("Expected delete action, got %v", operation.Action)
				}

				payloadSlice, ok := operation.Payload.([]interface{})
				if !ok {
					t.Error("Delete payload should be a slice")
				} else if len(payloadSlice) != 2 {
					t.Errorf("Expected 2 delete items, got %v", len(payloadSlice))
				}
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": true,
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

	entity := MockEntity{}
	ids := []string{"id1", "id2"}

	resp, err := client.Delete(ctx, entity, ids)
	if err != nil {
		t.Errorf("Delete() error = %v", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Delete() status = %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestClient_Delete_EmptyIDs(t *testing.T) {
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

		if r.URL.Path == "/api/_action/sync" && r.Method == "POST" {
			var payload map[string]SyncOperation
			err := json.NewDecoder(r.Body).Decode(&payload)
			if err != nil {
				t.Errorf("Failed to decode payload: %v", err)
			}

			for _, operation := range payload {
				payloadSlice, ok := operation.Payload.([]interface{})
				if !ok {
					t.Error("Payload should be a slice")
				} else if len(payloadSlice) != 0 {
					t.Errorf("Expected empty payload, got %v items", len(payloadSlice))
				}
			}

			w.WriteHeader(http.StatusOK)
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

	entity := MockEntity{}
	ids := []string{}

	resp, err := client.Delete(ctx, entity, ids)
	if err != nil {
		t.Errorf("Delete() with empty IDs error = %v", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Delete() status = %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestDeleteEntity_Structure(t *testing.T) {
	entity := deleteEntity{ID: "test-id"}

	data, err := json.Marshal(entity)
	if err != nil {
		t.Fatalf("Failed to marshal deleteEntity: %v", err)
	}

	var unmarshaled deleteEntity
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal deleteEntity: %v", err)
	}

	if entity.ID != unmarshaled.ID {
		t.Errorf("ID not preserved: got %v, want %v", unmarshaled.ID, entity.ID)
	}
}

func TestUpsert_ReflectionValidation(t *testing.T) {
	entities := []MockEntity{{ID: "1"}}

	if reflect.ValueOf(entities).Kind() != reflect.Slice {
		t.Error("Should detect slice type correctly")
	}

	singleEntity := MockEntity{ID: "1"}
	if reflect.ValueOf(singleEntity).Kind() == reflect.Slice {
		t.Error("Should not detect non-slice as slice")
	}
}

func TestClient_Sync_ErrorResponse(t *testing.T) {
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

		if r.URL.Path == "/api/_action/sync" && r.Method == "POST" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"errors": []map[string]interface{}{
					{
						"status": "400",
						"code":   "SYNC_ERROR",
						"detail": "Sync operation failed",
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
	operations := map[string]SyncOperation{
		"test-op": {
			Entity: "product",
			Action: "upsert",
			Payload: []map[string]interface{}{
				{"name": "Test Product"},
			},
		},
	}

	_, err = client.Sync(ctx, operations)
	if err == nil {
		t.Error("Sync() should return error when API returns error")
	}
}

func TestClient_Upsert_EmptySlice(t *testing.T) {
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

	client, err := NewClient(context.Background(), server.URL, NewIntegrationCredentials("test", "test", []string{"write"}), nil)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := NewAPIContext(context.Background())

	var emptyEntities []MockEntity
	_, err = client.Upsert(ctx, emptyEntities)
	if err == nil {
		t.Error("Upsert() should return error for empty slice")
	}
}

func TestClient_Sync_RequestCreationError(t *testing.T) {
	client := &Client{
		remote: "://invalid-url", // This will cause url.JoinPath to fail
	}

	ctx := NewAPIContext(context.Background())
	operations := map[string]SyncOperation{
		"test-op": {
			Entity: "product",
			Action: "upsert",
			Payload: []map[string]interface{}{
				{"name": "Test Product"},
			},
		},
	}

	_, err := client.Sync(ctx, operations)
	if err == nil {
		t.Error("Sync() should return error when NewRequest fails")
	}
}

func TestClient_Upsert_NotASlice(t *testing.T) {
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
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client, err := NewClient(context.Background(), server.URL, NewIntegrationCredentials("test", "test", []string{"write"}), nil)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := NewAPIContext(context.Background())

	singleEntity := MockEntity{ID: "1", Name: "Test"}
	_, err = client.Upsert(ctx, singleEntity)
	if err == nil {
		t.Error("Upsert() should return error for non-slice entity")
	}
	if err.Error() != "entity is not a slice" {
		t.Errorf("Upsert() error = %v, want 'entity is not a slice'", err)
	}
}
