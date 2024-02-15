package com

import (
	"net/http"

	"github.com/pkg/errors"
)

// Clear the shopware cache
func (c Client) Clear(ctx APIContext) (*http.Response, error) {
	r, err := c.NewRequest(ctx, http.MethodDelete, "/api/_action/cache", nil, nil)

	if err != nil {
		return nil, errors.Wrap(err, "failed to create request for cache clear")
	}

	return c.BareDo(ctx.Context, r)
}
