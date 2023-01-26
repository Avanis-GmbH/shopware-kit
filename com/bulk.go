package com

import "net/http"

func (c *Client) Sync(ctx ApiContext, payload map[string]SyncOperation) (*http.Response, error) {
	req, err := c.NewRequest(ctx, http.MethodPost, "/api/_action/sync", payload)

	if err != nil {
		return nil, err
	}

	return c.Do(ctx.Context, req, nil)
}

type SyncOperation struct {
	Entity  string      `json:"entity"`
	Action  string      `json:"action"`
	Payload interface{} `json:"payload"`
}
