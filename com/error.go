package com

import "fmt"

type APIError struct {
	StatusCode int
	Response   ErrorResponse
	Raw        []byte
}

type ErrorResponse struct {
	Errors []struct {
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

	Extensions []interface{} `json:"extensions"`
}

func (e APIError) Error() string {
	str := fmt.Sprintf("%d - errors: \n", e.StatusCode)
	for _, err := range e.Response.Errors {
		str += fmt.Sprintf("\t: %s -> %s => %s \n", err.Code, err.Detail, err.Source.Pointer)
	}

	return str
}
