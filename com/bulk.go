package com

import (
	"net/http"
	"reflect"

	"github.com/pkg/errors"
)

// Sync Operations for bulk request
// (https://shopware.stoplight.io/docs/admin-api/0612cb5d960ef-bulk-edit-entities)
type SyncOperation struct {
	Entity  string      `json:"entity"`
	Action  string      `json:"action"`
	Payload interface{} `json:"payload"`
}

type deleteEntity struct {
	Id string `json:"id"`
}

// Starts a sync process for the list of provided actions.
// This can be inserts, upserts, updates and deletes on different entities.
func (c *Client) Sync(ctx ApiContext, payload map[string]SyncOperation) (*http.Response, error) {
	req, err := c.NewRequest(ctx, http.MethodPost, "/api/_action/sync", payload)

	if err != nil {
		return nil, errors.Wrapf(err, "failed to create request for sync operation")
	}

	return c.Do(ctx.Context, req, nil)
}

// Inserts or updates the provided entities in bulk mode
// Provided entities must be a slice
// Uses underlying Sync method to perform the request
func (c *Client) Upsert(ctx ApiContext, entities interface{}) (*http.Response, error) {
	if c.GetSegmentSnakeCase(entities) == "unknown" {
		return nil, errors.New("unknown entity")
	}

	if reflect.ValueOf(entities).Kind() != reflect.Slice {
		return nil, errors.New("entity is not a slice")
	}

	return c.Sync(ctx, map[string]SyncOperation{c.GetSegmentSnakeCase(entities): {
		Entity:  c.GetSegmentSnakeCase(entities),
		Action:  "upsert",
		Payload: entities,
	}})
}

// Deletes the provided entities in bulk mode
// Uses underlying Sync method to perform the request
func (c *Client) Delete(ctx ApiContext, entity interface{}, ids []string) (*http.Response, error) {
	payload := make([]deleteEntity, 0)

	for _, id := range ids {
		payload = append(payload, deleteEntity{Id: id})
	}

	return c.Sync(ctx, map[string]SyncOperation{c.GetSegmentSnakeCase(entity): {
		Entity:  c.GetSegmentSnakeCase(entity),
		Action:  "delete",
		Payload: payload,
	}})
}
