package com

import (
	"net/http"

	"github.com/pkg/errors"
)

func (c Client) Clear(ctx ApiContext) (*http.Response, error) {
	r, err := c.NewRequest(ctx, http.MethodDelete, "/api/_action/cache", nil)

	if err != nil {
		return nil, errors.Wrap(err, "failed to create request for cache clear")
	}

	return c.BareDo(ctx.Context, r)
}
