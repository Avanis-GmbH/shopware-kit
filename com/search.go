package com

import (
	"fmt"
	"net/url"
	"reflect"
)

func (c *Client) Search(ctx ApiContext, criteria Criteria, v interface{}) ([]*interface{}, error) {
	url, err := url.JoinPath("/api/search", c.getSegment(v))
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest(ctx, "POST", url, criteria)
	if err != nil {
		return nil, err
	}

	s := reflect.MakeSlice(reflect.SliceOf(reflect.PtrTo(reflect.TypeOf(v))), 0, 0).Interface()
	data, ok := s.([]*interface{})
	if !ok {
		return nil, fmt.Errorf("failed to convert interface to []*interface{}")
	}

	_, err = c.Do(ctx.Context, req, data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (c *Client) SearchAll(ctx ApiContext, criteria Criteria, v interface{}) ([]*interface{}, error) {
	if criteria.Limit == 0 {
		criteria.Limit = 50
	}

	if criteria.Page == 0 {
		criteria.Page = 1
	}

	result := make([]*interface{}, 0)
	for {
		criteria.Page++

		data, err := c.Search(ctx, criteria, v)
		if err != nil {
			return nil, err
		}

		if len(data) == 0 {
			break
		}

		result = append(result, data...)
	}

	return result, nil
}

func (c *Client) SearchIds(ctx ApiContext, criteria Criteria, v interface{}) (*SearchIdsResponse, error) {
	url, err := url.JoinPath("/api/search-ids", c.getSegment(v))
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest(ctx, "POST", url, criteria)
	if err != nil {
		return nil, err
	}

	data := &SearchIdsResponse{}
	_, err = c.Do(ctx.Context, req, data)
	if err != nil {
		return data, err
	}

	return data, nil
}

type SearchIdsResponse struct {
	Total int      `json:"total"`
	Data  []string `json:"data"`
}
