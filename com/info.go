package com

import (
	"fmt"
	"net/http"
)

// InfoResponse is the response of the info request to the shopware api
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

// infoResponseBundle is used by the InfoResponse struct
type infoResponseBundle struct {
	CSS []string `json:"css"`
	Js  []string `json:"js"`
}

// IsCloudShop checks if the shop is a cloud shop
func (r InfoResponse) IsCloudShop() bool {
	_, ok := r.Bundles["SaasRufus"]

	return ok
}

// Info returns the info of the shopware api
func (c *Client) Info(ctx APIContext) (*InfoResponse, *http.Response, error) {
	r, err := c.NewRequest(ctx, http.MethodGet, "/api/_info/config", nil, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create request for info: %w", err)
	}

	var info *InfoResponse
	resp, err := c.Do(ctx.Context, r, &info)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to execute request for info: %w", err)
	}

	return info, resp, err
}
