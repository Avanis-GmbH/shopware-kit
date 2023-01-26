package com

import (
	"net/http"
	"reflect"

	"github.com/pkg/errors"
)

type SyncOperation struct {
	Entity  string      `json:"entity"`
	Action  string      `json:"action"`
	Payload interface{} `json:"payload"`
}

type DeleteEntity struct {
	Id string `json:"id"`
}

func (c *Client) Sync(ctx ApiContext, payload map[string]SyncOperation) (*http.Response, error) {
	req, err := c.NewRequest(ctx, http.MethodPost, "/api/_action/sync", payload)

	if err != nil {
		return nil, errors.Wrapf(err, "failed to create request for sync operation")
	}

	return c.Do(ctx.Context, req, nil)
}

func (c *Client) Upsert(ctx ApiContext, entity interface{}) (*http.Response, error) {
	if c.getSegment(entity) == "unknown" {
		return nil, errors.New("unknown entity")
	}

	if reflect.ValueOf(entity).Kind() != reflect.Slice {
		return nil, errors.New("entity is not a slice")
	}

	return c.Sync(ctx, map[string]SyncOperation{c.getSegment(entity): {
		Entity:  c.getSegment(entity),
		Action:  "upsert",
		Payload: entity,
	}})
}

func (c *Client) Delete(ctx ApiContext, entity interface{}, ids []string) (*http.Response, error) {
	payload := make([]DeleteEntity, 0)

	for _, id := range ids {
		payload = append(payload, DeleteEntity{Id: id})
	}

	return c.Sync(ctx, map[string]SyncOperation{c.getSegment(entity): {
		Entity:  c.getSegment(entity),
		Action:  "delete",
		Payload: payload,
	}})
}
