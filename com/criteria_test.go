package com

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestCriteria_JSONSerialization(t *testing.T) {
	criteria := Criteria{
		Includes: map[string][]string{
			"product": {"id", "name"},
		},
		Page:           1,
		Limit:          10,
		IDs:            []string{"id1", "id2"},
		Term:           "search term",
		TotalCountMode: TotalCountModeExact,
	}

	data, err := json.Marshal(criteria)
	if err != nil {
		t.Fatalf("Failed to marshal criteria: %v", err)
	}

	var unmarshaled Criteria
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal criteria: %v", err)
	}

	if !reflect.DeepEqual(criteria.Includes, unmarshaled.Includes) {
		t.Errorf("Includes not preserved: got %v, want %v", unmarshaled.Includes, criteria.Includes)
	}

	if criteria.Page != unmarshaled.Page {
		t.Errorf("Page not preserved: got %v, want %v", unmarshaled.Page, criteria.Page)
	}

	if criteria.Limit != unmarshaled.Limit {
		t.Errorf("Limit not preserved: got %v, want %v", unmarshaled.Limit, criteria.Limit)
	}

	if criteria.Term != unmarshaled.Term {
		t.Errorf("Term not preserved: got %v, want %v", unmarshaled.Term, criteria.Term)
	}
}

func TestCriteriaFilter(t *testing.T) {
	filter := CriteriaFilter{
		Type:     SearchFilterTypeEquals,
		Operator: "AND",
		Queries:  []interface{}{"value1", "value2"},
	}

	data, err := json.Marshal(filter)
	if err != nil {
		t.Fatalf("Failed to marshal filter: %v", err)
	}

	var unmarshaled CriteriaFilter
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal filter: %v", err)
	}

	if filter.Type != unmarshaled.Type {
		t.Errorf("Type not preserved: got %v, want %v", unmarshaled.Type, filter.Type)
	}

	if filter.Operator != unmarshaled.Operator {
		t.Errorf("Operator not preserved: got %v, want %v", unmarshaled.Operator, filter.Operator)
	}
}

func TestCriteriaFilterQuery(t *testing.T) {
	filterQuery := CriteriaFilterQuery{
		CriteriaFilter: CriteriaFilter{
			Type: SearchFilterTypeEquals,
		},
		Field: "name",
		Value: "test value",
	}

	data, err := json.Marshal(filterQuery)
	if err != nil {
		t.Fatalf("Failed to marshal filter query: %v", err)
	}

	var unmarshaled CriteriaFilterQuery
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal filter query: %v", err)
	}

	if filterQuery.Field != unmarshaled.Field {
		t.Errorf("Field not preserved: got %v, want %v", unmarshaled.Field, filterQuery.Field)
	}

	if filterQuery.Value != unmarshaled.Value {
		t.Errorf("Value not preserved: got %v, want %v", unmarshaled.Value, filterQuery.Value)
	}

	if filterQuery.Type != unmarshaled.Type {
		t.Errorf("Type not preserved: got %v, want %v", unmarshaled.Type, filterQuery.Type)
	}
}

func TestCriteriaQuery(t *testing.T) {
	query := CriteriaQuery{
		Score: 100,
		Query: CriteriaFilterQuery{
			CriteriaFilter: CriteriaFilter{
				Type: SearchFilterTypeContains,
			},
			Field: "description",
			Value: "search term",
		},
	}

	data, err := json.Marshal(query)
	if err != nil {
		t.Fatalf("Failed to marshal criteria query: %v", err)
	}

	var unmarshaled CriteriaQuery
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal criteria query: %v", err)
	}

	if query.Score != unmarshaled.Score {
		t.Errorf("Score not preserved: got %v, want %v", unmarshaled.Score, query.Score)
	}

	if query.Query.Field != unmarshaled.Query.Field {
		t.Errorf("Query.Field not preserved: got %v, want %v", unmarshaled.Query.Field, query.Query.Field)
	}

	if query.Query.Value != unmarshaled.Query.Value {
		t.Errorf("Query.Value not preserved: got %v, want %v", unmarshaled.Query.Value, query.Query.Value)
	}
}

func TestCriteriaSort(t *testing.T) {
	sort := CriteriaSort{
		Direction:      SearchSortDirectionAscending,
		Field:          "name",
		NaturalSorting: true,
	}

	data, err := json.Marshal(sort)
	if err != nil {
		t.Fatalf("Failed to marshal sort: %v", err)
	}

	var unmarshaled CriteriaSort
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal sort: %v", err)
	}

	if sort.Direction != unmarshaled.Direction {
		t.Errorf("Direction not preserved: got %v, want %v", unmarshaled.Direction, sort.Direction)
	}

	if sort.Field != unmarshaled.Field {
		t.Errorf("Field not preserved: got %v, want %v", unmarshaled.Field, sort.Field)
	}

	if sort.NaturalSorting != unmarshaled.NaturalSorting {
		t.Errorf("NaturalSorting not preserved: got %v, want %v", unmarshaled.NaturalSorting, sort.NaturalSorting)
	}
}

func TestCriteria_WithAssociations(t *testing.T) {
	criteria := Criteria{
		Page:  1,
		Limit: 10,
		Associations: map[string]Criteria{
			"manufacturer": {
				Includes: map[string][]string{
					"manufacturer": {"id", "name"},
				},
			},
			"categories": {
				Limit: 5,
			},
		},
	}

	data, err := json.Marshal(criteria)
	if err != nil {
		t.Fatalf("Failed to marshal criteria with associations: %v", err)
	}

	var unmarshaled Criteria
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal criteria with associations: %v", err)
	}

	if len(unmarshaled.Associations) != 2 {
		t.Errorf("Associations length: got %v, want 2", len(unmarshaled.Associations))
	}

	manufacturerAssoc, exists := unmarshaled.Associations["manufacturer"]
	if !exists {
		t.Error("manufacturer association not found")
	} else {
		if len(manufacturerAssoc.Includes) != 1 {
			t.Errorf("manufacturer association includes length: got %v, want 1", len(manufacturerAssoc.Includes))
		}
	}

	categoriesAssoc, exists := unmarshaled.Associations["categories"]
	if !exists {
		t.Error("categories association not found")
	} else {
		if categoriesAssoc.Limit != 5 {
			t.Errorf("categories association limit: got %v, want 5", categoriesAssoc.Limit)
		}
	}
}

func TestConstants(t *testing.T) {
	if TotalCountModeDefault != 1 {
		t.Errorf("TotalCountModeDefault = %v, want 1", TotalCountModeDefault)
	}
	if TotalCountModeExact != 2 {
		t.Errorf("TotalCountModeExact = %v, want 2", TotalCountModeExact)
	}
	if TotalCountModeNextPage != 3 {
		t.Errorf("TotalCountModeNextPage = %v, want 3", TotalCountModeNextPage)
	}

	filterTypes := []SearchFilterType{
		SearchFilterTypeEquals,
		SearchFilterTypeEqualsAny,
		SearchFilterTypeContains,
		SearchFilterTypeRange,
		SearchFilterTypeNot,
		SearchFilterTypeMulti,
		SearchFilterTypePrefix,
		SearchFilterTypeSuffix,
	}

	expectedTypes := []string{
		"equals", "equalsAny", "contains", "range", "not", "multi", "prefix", "suffix",
	}

	for i, filterType := range filterTypes {
		if string(filterType) != expectedTypes[i] {
			t.Errorf("SearchFilterType[%d] = %v, want %v", i, filterType, expectedTypes[i])
		}
	}

	if SearchSortDirectionAscending != "ASC" {
		t.Errorf("SearchSortDirectionAscending = %v, want ASC", SearchSortDirectionAscending)
	}
	if SearchSortDirectionDescending != "DESC" {
		t.Errorf("SearchSortDirectionDescending = %v, want DESC", SearchSortDirectionDescending)
	}
}

func TestCriteria_EmptyValues(t *testing.T) {
	criteria := Criteria{}

	data, err := json.Marshal(criteria)
	if err != nil {
		t.Fatalf("Failed to marshal empty criteria: %v", err)
	}

	var jsonMap map[string]interface{}
	err = json.Unmarshal(data, &jsonMap)
	if err != nil {
		t.Fatalf("Failed to unmarshal to map: %v", err)
	}

	// Only non-zero values should be present
	for key := range jsonMap {
		t.Errorf("Empty criteria should not contain key: %v", key)
	}
}

func TestCriteria_WithFiltersAndSort(t *testing.T) {
	criteria := Criteria{
		Filter: []interface{}{
			CriteriaFilterQuery{
				CriteriaFilter: CriteriaFilter{
					Type: SearchFilterTypeEquals,
				},
				Field: "active",
				Value: true,
			},
		},
		PostFilter: []interface{}{
			CriteriaFilterQuery{
				CriteriaFilter: CriteriaFilter{
					Type: SearchFilterTypeRange,
				},
				Field: "price",
				Value: map[string]interface{}{"gte": 10, "lte": 100},
			},
		},
		Sort: []CriteriaSort{
			{
				Direction:      SearchSortDirectionDescending,
				Field:          "createdAt",
				NaturalSorting: false,
			},
		},
	}

	data, err := json.Marshal(criteria)
	if err != nil {
		t.Fatalf("Failed to marshal criteria: %v", err)
	}

	var unmarshaled Criteria
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal criteria: %v", err)
	}

	if len(unmarshaled.Filter) != 1 {
		t.Errorf("Filter length: got %v, want 1", len(unmarshaled.Filter))
	}

	if len(unmarshaled.PostFilter) != 1 {
		t.Errorf("PostFilter length: got %v, want 1", len(unmarshaled.PostFilter))
	}

	if len(unmarshaled.Sort) != 1 {
		t.Errorf("Sort length: got %v, want 1", len(unmarshaled.Sort))
	}

	if unmarshaled.Sort[0].Field != "createdAt" {
		t.Errorf("Sort field: got %v, want createdAt", unmarshaled.Sort[0].Field)
	}
}
