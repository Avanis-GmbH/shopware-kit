package com

import (
	"net/url"

	"github.com/Avanis-GmbH/SUC/model"
)

func (c *Client) Search(ctx ApiContext, criteria Criteria, v model.Collection) error {
	url, err := url.JoinPath("/api/search", c.getSegment(v))
	if err != nil {
		return err
	}

	req, err := c.NewRequest(ctx, "POST", url, criteria)
	if err != nil {
		return err
	}

	_, err = c.Do(ctx.Context, req, v)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) SearchAll(ctx ApiContext, criteria Criteria, v model.Collection) error {
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
			return err
		}

		if len(v.GetData()) == 0 {
			break
		}

		result = append(result, v.GetData()...)
	}

	v.SetTotal(int64(len(result)))
	v.SetData(result)
	return nil
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
