package com

import "fmt"

type APIError struct {
	StatusCode int
	Response   ErrorResponse
	Raw        []byte
}

type ErrorResponse struct {
	Success bool `json:"success"`
	Data    map[string]struct {
		Result []struct {
			Entities []interface{} `json:"entities"`
			Errors   []struct {
				Status string `json:"status"`
				Code   string `json:"code"`
				Title  string `json:"title"`
				Detail string `json:"detail"`
				Source struct {
					Pointer string `json:"pointer"`
				} `json:"source"`
				Meta struct {
					Parameters interface{} `json:"parameters"`
				} `json:"meta"`
			} `json:"errors"`
		} `json:"result"`
		Extensions []interface{} `json:"extensions"`
	} `json:"data"`
}

func (e APIError) Error() string {
	str := fmt.Sprintf("%d - errrors: \n", e.StatusCode)
	for subject, data := range e.Response.Data {
		for _, result := range data.Result {
			for _, err := range result.Errors {
				str += fmt.Sprintf("\t%s: %s -> %s => %s \n", subject, err.Code, err.Detail, err.Source.Pointer)
			}
		}
	}

	return str
}
