package com

import (
	"reflect"
	"testing"
)

type MockEntity struct {
	ID   string
	Name string
}

type MockEntityCollection struct {
	EntityCollection
	Data []MockEntity
}

func TestClient_GetSegment(t *testing.T) {
	client := &Client{}

	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{
			name:     "simple struct",
			input:    MockEntity{},
			expected: "mock-entity",
		},
		{
			name:     "pointer to struct",
			input:    &MockEntity{},
			expected: "mock-entity",
		},
		{
			name:     "collection struct",
			input:    MockEntityCollection{},
			expected: "mock-entity",
		},
		{
			name:     "pointer to collection",
			input:    &MockEntityCollection{},
			expected: "mock-entity",
		},
		{
			name:     "slice with elements",
			input:    []MockEntity{{ID: "1"}},
			expected: "mock-entity",
		},
		{
			name:     "empty slice",
			input:    []MockEntity{},
			expected: "unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := client.GetSegment(tt.input)
			if result != tt.expected {
				t.Errorf("GetSegment(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestClient_GetSegmentSnakeCase(t *testing.T) {
	client := &Client{}

	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{
			name:     "simple struct",
			input:    MockEntity{},
			expected: "mock_entity",
		},
		{
			name:     "collection struct",
			input:    MockEntityCollection{},
			expected: "mock_entity",
		},
		{
			name:     "pointer to struct",
			input:    &MockEntity{},
			expected: "mock_entity",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := client.GetSegmentSnakeCase(tt.input)
			if result != tt.expected {
				t.Errorf("GetSegmentSnakeCase(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestGetSegment_ReflectionEdgeCases(t *testing.T) {
	client := &Client{}

	type ProductManufacturerCollection struct{}
	type UserRolePermissionCollection struct{}

	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{
			name:     "complex name with collection suffix",
			input:    ProductManufacturerCollection{},
			expected: "product-manufacturer",
		},
		{
			name:     "multiple words with collection suffix",
			input:    UserRolePermissionCollection{},
			expected: "user-role-permission",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := client.GetSegment(tt.input)
			if result != tt.expected {
				t.Errorf("GetSegment(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestGetSegment_WithSliceOfPointers(t *testing.T) {
	client := &Client{}

	entities := []*MockEntity{{ID: "1"}}
	result := client.GetSegment(entities)
	expected := "mock-entity"

	if result != expected {
		t.Errorf("GetSegment(slice of pointers) = %v, want %v", result, expected)
	}
}

func TestGetSegment_RecursiveSliceHandling(t *testing.T) {
	client := &Client{}

	entities := []MockEntity{{ID: "1", Name: "test"}}
	result := client.GetSegment(entities)

	if result != "mock-entity" {
		t.Errorf("GetSegment should handle slice recursively, got %v", result)
	}
}

func TestReflection_TypeChecking(t *testing.T) {
	client := &Client{}
	entity := MockEntity{}

	if reflect.TypeOf(&entity).Kind() != reflect.Ptr {
		t.Error("Should detect pointer type correctly")
	}

	slice := []MockEntity{}
	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		t.Error("Should detect slice type correctly")
	}

	if client.GetSegment(&entity) != "mock-entity" {
		t.Error("GetSegment should handle pointer correctly")
	}

	if client.GetSegment(slice) != "unknown" {
		t.Error("GetSegment should handle empty slice correctly")
	}
}
