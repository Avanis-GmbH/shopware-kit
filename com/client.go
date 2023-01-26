package com

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

type Client struct {
	remote          string
	client          *http.Client
	ctx             context.Context
	credentials     OAuthCredentials
	ResponseHandler func(resp *http.Response) error
}

func NewClient(ctx context.Context, shopURL string, credentials OAuthCredentials, httpClient *http.Client) (*Client, error) {
	c := &Client{ctx: ctx, remote: shopURL, credentials: credentials, client: httpClient}

	err := c.authorize()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) authorize() error {
	if c.client != nil {
		c.ctx = context.WithValue(c.ctx, oauth2.HTTPClient, c.client)
	}

	url, err := url.JoinPath(c.remote, "/api/oauth/token")
	if err != nil {
		return err
	}

	tokenSrc, err := c.credentials.GetTokenSource(c.ctx, url)
	if err != nil {
		return err
	}

	c.client = oauth2.NewClient(c.ctx, tokenSrc)
	return nil
}

func (c *Client) BareDo(ctx context.Context, req *http.Request) (*http.Response, error) {
	if ctx == nil {
		return nil, errors.New("context is nil")
	}

	resp, err := c.client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, errors.Wrapf(err, "failed to execute request")
	}

	err = c.checkResponse(resp)
	return resp, err
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.BareDo(ctx, req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	return resp, errors.Wrap(err, "failed to decode response body")
}

func (c *Client) NewRequest(context ApiContext, method, urlStr string, body interface{}) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := c.NewRawRequest(context, method, urlStr, buf)

	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func (c *Client) NewRawRequest(context ApiContext, method, urlStr string, body io.Reader) (*http.Request, error) {
	url, err := url.JoinPath(c.remote, urlStr)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(context.Context, method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("sw-language-id", context.LanguageId)
	req.Header.Set("sw-version-id", context.VersionId)

	if context.SkipFlows {
		req.Header.Set("sw-skip-trigger-flow", "1")
	}

	req.Header.Set("Accept", "application/json")

	return req, nil
}

func (c *Client) checkResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	if c.ResponseHandler != nil {
		return c.ResponseHandler(r)
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("parsing json response failed: %w", err)
	}

	return fmt.Errorf("request failed: %s", string(data))
}

type ApiContext struct {
	Context    context.Context
	LanguageId string
	VersionId  string
	SkipFlows  bool
}

func NewApiContext(ctx context.Context) ApiContext {
	return ApiContext{
		Context:    ctx,
		LanguageId: "2fbb5fe2e29a4d70aa5854ce7ce3e20b",
		VersionId:  "0fa91ce3e96a4bc2be4bd9ce752c3425",
		SkipFlows:  false,
	}
}
