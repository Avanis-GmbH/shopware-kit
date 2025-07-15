package com

import (
	"encoding/json"
	"testing"
)

func TestAPIError_Error(t *testing.T) {
	apiError := APIError{
		StatusCode: 400,
		Response: ErrorResponse{
			Errors: []struct {
				Status string `json:"status"`
				Code   string `json:"code"`
				Title  string `json:"title"`
				Detail string `json:"detail"`
				Source struct {
					Pointer string `json:"pointer"`
				} `json:"source"`
				Meta struct {
					Parameters interface{} `json:"parameters"`
				} `json:"meta"`
			}{
				{
					Status: "400",
					Code:   "INVALID_REQUEST",
					Title:  "Bad Request",
					Detail: "The request is invalid",
					Source: struct {
						Pointer string `json:"pointer"`
					}{
						Pointer: "/data/attributes/name",
					},
					Meta: struct {
						Parameters interface{} `json:"parameters"`
					}{
						Parameters: map[string]string{"field": "name"},
					},
				},
				{
					Status: "400",
					Code:   "VALIDATION_ERROR",
					Title:  "Validation Error",
					Detail: "Name is required",
					Source: struct {
						Pointer string `json:"pointer"`
					}{
						Pointer: "/data/attributes/name",
					},
				},
			},
		},
		Raw: []byte(`{"errors": [{"status": "400", "code": "INVALID_REQUEST"}]}`),
	}

	errorString := apiError.Error()

	// Check that error message contains status code
	if errorString == "" {
		t.Error("Error() should return non-empty string")
	}

	// Check that it contains the status code
	if !contains(errorString, "400") {
		t.Error("Error() should contain status code")
	}

	// Check that it contains error codes
	if !contains(errorString, "INVALID_REQUEST") {
		t.Error("Error() should contain error code INVALID_REQUEST")
	}

	if !contains(errorString, "VALIDATION_ERROR") {
		t.Error("Error() should contain error code VALIDATION_ERROR")
	}

	// Check that it contains error details
	if !contains(errorString, "The request is invalid") {
		t.Error("Error() should contain error detail")
	}

	// Check that it contains source pointers
	if !contains(errorString, "/data/attributes/name") {
		t.Error("Error() should contain source pointer")
	}
}

func TestAPIError_EmptyErrors(t *testing.T) {
	apiError := APIError{
		StatusCode: 500,
		Response: ErrorResponse{
			Errors: []struct {
				Status string `json:"status"`
				Code   string `json:"code"`
				Title  string `json:"title"`
				Detail string `json:"detail"`
				Source struct {
					Pointer string `json:"pointer"`
				} `json:"source"`
				Meta struct {
					Parameters interface{} `json:"parameters"`
				} `json:"meta"`
			}{},
		},
		Raw: []byte(`{"errors": []}`),
	}

	errorString := apiError.Error()

	// Should still contain status code even with empty errors
	if !contains(errorString, "500") {
		t.Error("Error() should contain status code even with empty errors")
	}
}

func TestErrorResponse_JSONSerialization(t *testing.T) {
	errorResp := ErrorResponse{
		Errors: []struct {
			Status string `json:"status"`
			Code   string `json:"code"`
			Title  string `json:"title"`
			Detail string `json:"detail"`
			Source struct {
				Pointer string `json:"pointer"`
			} `json:"source"`
			Meta struct {
				Parameters interface{} `json:"parameters"`
			} `json:"meta"`
		}{
			{
				Status: "400",
				Code:   "VALIDATION_ERROR",
				Title:  "Validation Failed",
				Detail: "The name field is required",
				Source: struct {
					Pointer string `json:"pointer"`
				}{
					Pointer: "/data/attributes/name",
				},
				Meta: struct {
					Parameters interface{} `json:"parameters"`
				}{
					Parameters: map[string]interface{}{
						"field": "name",
						"rule":  "required",
					},
				},
			},
		},
		Extensions: []interface{}{
			map[string]string{
				"type": "validation",
			},
		},
	}

	data, err := json.Marshal(errorResp)
	if err != nil {
		t.Fatalf("Failed to marshal ErrorResponse: %v", err)
	}

	var unmarshaled ErrorResponse
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal ErrorResponse: %v", err)
	}

	if len(unmarshaled.Errors) != 1 {
		t.Errorf("Errors length: got %v, want 1", len(unmarshaled.Errors))
	}

	if unmarshaled.Errors[0].Code != "VALIDATION_ERROR" {
		t.Errorf("Error code: got %v, want VALIDATION_ERROR", unmarshaled.Errors[0].Code)
	}

	if unmarshaled.Errors[0].Detail != "The name field is required" {
		t.Errorf("Error detail: got %v, want 'The name field is required'", unmarshaled.Errors[0].Detail)
	}

	if unmarshaled.Errors[0].Source.Pointer != "/data/attributes/name" {
		t.Errorf("Error source pointer: got %v, want '/data/attributes/name'", unmarshaled.Errors[0].Source.Pointer)
	}

	if len(unmarshaled.Extensions) != 1 {
		t.Errorf("Extensions length: got %v, want 1", len(unmarshaled.Extensions))
	}
}

func TestAPIError_WithRawBytes(t *testing.T) {
	rawJSON := `{
		"errors": [
			{
				"status": "404",
				"code": "NOT_FOUND",
				"title": "Not Found",
				"detail": "The requested resource was not found",
				"source": {
					"pointer": "/api/product/123"
				}
			}
		]
	}`

	apiError := APIError{
		StatusCode: 404,
		Raw:        []byte(rawJSON),
	}

	// Unmarshal the raw bytes into the Response
	err := json.Unmarshal(apiError.Raw, &apiError.Response)
	if err != nil {
		t.Fatalf("Failed to unmarshal raw JSON: %v", err)
	}

	if len(apiError.Response.Errors) != 1 {
		t.Errorf("Errors length: got %v, want 1", len(apiError.Response.Errors))
	}

	if apiError.Response.Errors[0].Code != "NOT_FOUND" {
		t.Errorf("Error code: got %v, want NOT_FOUND", apiError.Response.Errors[0].Code)
	}

	errorString := apiError.Error()
	if !contains(errorString, "404") {
		t.Error("Error() should contain status code 404")
	}

	if !contains(errorString, "NOT_FOUND") {
		t.Error("Error() should contain error code NOT_FOUND")
	}
}

func TestAPIError_MultipleErrors(t *testing.T) {
	apiError := APIError{
		StatusCode: 422,
		Response: ErrorResponse{
			Errors: []struct {
				Status string `json:"status"`
				Code   string `json:"code"`
				Title  string `json:"title"`
				Detail string `json:"detail"`
				Source struct {
					Pointer string `json:"pointer"`
				} `json:"source"`
				Meta struct {
					Parameters interface{} `json:"parameters"`
				} `json:"meta"`
			}{
				{
					Status: "422",
					Code:   "VALIDATION_ERROR",
					Title:  "Validation Error",
					Detail: "Name is required",
					Source: struct {
						Pointer string `json:"pointer"`
					}{
						Pointer: "/data/attributes/name",
					},
				},
				{
					Status: "422",
					Code:   "VALIDATION_ERROR",
					Title:  "Validation Error",
					Detail: "Price must be positive",
					Source: struct {
						Pointer string `json:"pointer"`
					}{
						Pointer: "/data/attributes/price",
					},
				},
			},
		},
	}

	errorString := apiError.Error()

	// Check that both errors are included
	if !contains(errorString, "Name is required") {
		t.Error("Error() should contain first validation error")
	}

	if !contains(errorString, "Price must be positive") {
		t.Error("Error() should contain second validation error")
	}

	if !contains(errorString, "/data/attributes/name") {
		t.Error("Error() should contain first source pointer")
	}

	if !contains(errorString, "/data/attributes/price") {
		t.Error("Error() should contain second source pointer")
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 || s[0:len(substr)] == substr ||
		(len(s) > len(substr) && (s[1:len(substr)+1] == substr || contains(s[1:], substr))))
}

// More efficient contains implementation
func containsSubstring(s, substr string) bool {
	if len(substr) == 0 {
		return true
	}
	if len(s) < len(substr) {
		return false
	}
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func TestAPIError_Interface(t *testing.T) {
	apiError := APIError{
		StatusCode: 400,
		Response: ErrorResponse{
			Errors: []struct {
				Status string `json:"status"`
				Code   string `json:"code"`
				Title  string `json:"title"`
				Detail string `json:"detail"`
				Source struct {
					Pointer string `json:"pointer"`
				} `json:"source"`
				Meta struct {
					Parameters interface{} `json:"parameters"`
				} `json:"meta"`
			}{
				{
					Status: "400",
					Code:   "TEST_ERROR",
					Detail: "Test error message",
				},
			},
		},
	}

	// Test that APIError implements the error interface
	var err error = apiError
	if err.Error() == "" {
		t.Error("APIError should implement error interface")
	}

	// Test pointer to APIError also implements error interface
	var errPtr error = &apiError
	if errPtr.Error() == "" {
		t.Error("*APIError should implement error interface")
	}
}
