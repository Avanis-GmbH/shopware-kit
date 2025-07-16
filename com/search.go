package com

import (
	"fmt"
	"net/url"
)

// Search performs a search request for the given criteria and collection
// The collection must be a pointer to a struct that implements the Collection interface
func (c *Client) Search(ctx APIContext, criteria Criteria, v Collection) error {
	url, err := url.JoinPath("/api/search", c.GetSegment(v))
	if err != nil {
		return fmt.Errorf("failed to join path for %T: %w", v, err)
	}

	req, err := c.NewRequest(ctx, "POST", url, nil, criteria)
	if err != nil {
		return fmt.Errorf("failed to create request for %T: %w", v, err)
	}

	resp, err := c.Do(ctx.Context, req, v)
	if err != nil {
		return fmt.Errorf("failed to do Search request for %T: %w", v, err)
	}
	defer resp.Body.Close()

	return nil
}

// SearchAll performs multiple search requests with a limited size until all results are fetched
func (c *Client) SearchAll(ctx APIContext, criteria Criteria, v Collection) error {
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
			return fmt.Errorf("failed to search for %T: %w", v, err)
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

// SearchIDs performs a search request for the given criteria and returns only the ids
func (c *Client) SearchIDs(ctx APIContext, criteria Criteria, v interface{}) (*SearchIDsResponse, error) {
	url, err := url.JoinPath("/api/search-ids", c.GetSegment(v))
	if err != nil {
		return nil, fmt.Errorf("failed to join path for %T: %w", v, err)
	}

	req, err := c.NewRequest(ctx, "POST", url, nil, criteria)
	if err != nil {
		return nil, fmt.Errorf("failed to create request for %T: %w", v, err)
	}

	data := &SearchIDsResponse{}
	resp, err := c.Do(ctx.Context, req, data)
	if err != nil {
		return data, fmt.Errorf("failed to do searchIds request for %T: %w", v, err)
	}
	defer resp.Body.Close()

	return data, nil
}

// SearchIDsResponse is the response for a searchIds request
type SearchIDsResponse struct {
	Total int      `json:"total"`
	Data  []string `json:"data"`
}
