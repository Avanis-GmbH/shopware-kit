package com

import (
	"net/url"

	"github.com/pkg/errors"
)

// Search performs a search request for the given criteria and collection
// The collection must be a pointer to a struct that implements the Collection interface
func (c *Client) Search(ctx ApiContext, criteria Criteria, v Collection) error {
	url, err := url.JoinPath("/api/search", c.GetSegment(v))
	if err != nil {
		return errors.Wrapf(err, "failed to join path for %T", v)
	}

	req, err := c.NewRequest(ctx, "POST", url, nil, criteria)
	if err != nil {
		return errors.Wrapf(err, "failed to create request for %T", v)
	}

	_, err = c.Do(ctx.Context, req, v)
	if err != nil {
		return errors.Wrapf(err, "failed to do Search request for %T", v)
	}

	return nil
}

// SearchAll performs multiple search requests with a limited size until all results are fetched
func (c *Client) SearchAll(ctx ApiContext, criteria Criteria, v Collection) error {
	if criteria.Limit == 0 {
		criteria.Limit = 50
	}

	if criteria.Page == 0 {
		criteria.Page = 1
	}

	result := make([]interface{}, 0)
	for {
		criteria.Page++

		err := c.Search(ctx, criteria, v)
		if err != nil {
			return errors.Wrapf(err, "failed to search for %T", v)
		}

		if len(v.getData()) == 0 {
			break
		}

		result = append(result, v.getData()...)
	}

	v.setTotal(int64(len(result)))
	v.setData(result)
	return nil
}

// SearchIds performs a search request for the given criteria and returns only the ids
func (c *Client) SearchIds(ctx ApiContext, criteria Criteria, v interface{}) (*SearchIdsResponse, error) {
	url, err := url.JoinPath("/api/search-ids", c.GetSegment(v))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to join path for %T", v)
	}

	req, err := c.NewRequest(ctx, "POST", url, nil, criteria)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create request for %T", v)
	}

	data := &SearchIdsResponse{}
	_, err = c.Do(ctx.Context, req, data)
	if err != nil {
		return data, errors.Wrapf(err, "failed to do searchIds request for %T", v)
	}

	return data, nil
}

// SearchIdsResponse is the response for a searchIds request
type SearchIdsResponse struct {
	Total int      `json:"total"`
	Data  []string `json:"data"`
}
