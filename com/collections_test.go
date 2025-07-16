package com

import (
	"testing"
)

func TestEntityCollection_SetAndGetTotal(t *testing.T) {
	collection := EntityCollection{}

	collection.setTotal(100)
	if collection.getTotal() != 100 {
		t.Errorf("getTotal() = %v, want 100", collection.getTotal())
	}

	if collection.Total != 100 {
		t.Errorf("Total = %v, want 100", collection.Total)
	}
}

func TestEntityCollection_SetAndGetAggregations(t *testing.T) {
	collection := EntityCollection{}

	aggregations := map[string]interface{}{
		"count": 50,
		"avg":   25.5,
	}

	collection.setAggregations(aggregations)
	retrieved := collection.getAggregations()

	if retrieved == nil {
		t.Error("getAggregations() returned nil")
	}

	if collection.Aggregations == nil {
		t.Error("Aggregations is nil")
	}
}

func TestEntityCollection_SetAndGetData(t *testing.T) {
	collection := EntityCollection{}

	data := []interface{}{
		map[string]string{"id": "1", "name": "item1"},
		map[string]string{"id": "2", "name": "item2"},
	}

	collection.setData(data)
	retrieved := collection.getData()

	if len(retrieved) != 2 {
		t.Errorf("getData() length = %v, want 2", len(retrieved))
	}

	if len(collection.Data) != 2 {
		t.Errorf("Data length = %v, want 2", len(collection.Data))
	}
}

func TestEntityCollection_Interface(t *testing.T) {
	var collection Collection = &EntityCollection{}

	collection.setTotal(10)
	if collection.getTotal() != 10 {
		t.Error("EntityCollection should implement Collection interface")
	}

	aggregations := map[string]int{"test": 1}
	collection.setAggregations(aggregations)
	if collection.getAggregations() == nil {
		t.Error("EntityCollection should implement Collection interface")
	}

	data := []interface{}{"test"}
	collection.setData(data)
	if len(collection.getData()) != 1 {
		t.Error("EntityCollection should implement Collection interface")
	}
}

func TestACLRoleCollection_Interface(t *testing.T) {
	var collection Collection = &ACLRoleCollection{}

	collection.setTotal(5)
	if collection.getTotal() != 5 {
		t.Error("ACLRoleCollection should implement Collection interface through EntityCollection")
	}
}

func TestSpecificCollections_Interface(t *testing.T) {
	collections := []Collection{
		&ACLRoleCollection{},
		&ACLUserRoleCollection{},
		&AppCollection{},
		&CategoryCollection{},
		&CountryCollection{},
		&CurrencyCollection{},
		&CustomerCollection{},
	}

	for i, collection := range collections {
		collection.setTotal(int64(i + 1))
		if collection.getTotal() != int64(i+1) {
			t.Errorf("Collection %T should implement Collection interface", collection)
		}
	}
}

func TestEntityCollection_EmptyData(t *testing.T) {
	collection := EntityCollection{}

	emptyData := []interface{}{}
	collection.setData(emptyData)

	if len(collection.getData()) != 0 {
		t.Errorf("getData() with empty data should return length 0, got %v", len(collection.getData()))
	}

	if collection.getTotal() != 0 {
		t.Errorf("getTotal() with no total set should return 0, got %v", collection.getTotal())
	}
}

func TestEntityCollection_NilAggregations(t *testing.T) {
	collection := EntityCollection{}

	collection.setAggregations(nil)
	retrieved := collection.getAggregations()

	if retrieved != nil {
		t.Errorf("getAggregations() with nil should return nil, got %v", retrieved)
	}
}
