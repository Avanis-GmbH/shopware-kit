package com

import "net/http"

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
		return nil, err
	}

	return c.Do(ctx.Context, req, nil)
}

func (c *Client) Upsert(ctx ApiContext, entity []*interface{}) (*http.Response, error) {
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
