package com

import (
	"net/http"

	"github.com/pkg/errors"
)

type InfoResponse struct {
	Version         string `json:"version"`
	VersionRevision string `json:"versionRevision"`
	AdminWorker     struct {
		EnableAdminWorker bool     `json:"enableAdminWorker"`
		Transports        []string `json:"transports"`
	} `json:"adminWorker"`
	Bundles  map[string]infoResponseBundle `json:"bundles"`
	Settings struct {
		EnableURLFeature bool `json:"enableUrlFeature"`
	} `json:"settings"`
}

type infoResponseBundle struct {
	CSS []string `json:"css"`
	Js  []string `json:"js"`
}

func (r InfoResponse) IsCloudShop() bool {
	_, ok := r.Bundles["SaasRufus"]

	return ok
}

func (c *Client) Info(ctx ApiContext) (*InfoResponse, *http.Response, error) {
	r, err := c.NewRequest(ctx, http.MethodGet, "/api/_info/config", nil)
	if err != nil {
		return nil, nil, errors.Wrap(err, "could not create request")
	}

	var info *InfoResponse
	resp, err := c.Do(ctx.Context, r, &info)
	if err != nil {
		return nil, nil, err
	}

	return info, resp, err
}
