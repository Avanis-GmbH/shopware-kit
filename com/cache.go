package com

import (
	"fmt"
	"net/http"
)

// Clear the shopware cache
func (c Client) Clear(ctx APIContext) (*http.Response, error) {
	r, err := c.NewRequest(ctx, http.MethodDelete, "/api/_action/cache", nil, nil)

	if err != nil {
		return nil, fmt.Errorf("failed to create request for cache clear: %w", err)
	}

	return c.BareDo(ctx.Context, r)
}
