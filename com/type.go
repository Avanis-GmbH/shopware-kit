package com

import (
	"reflect"

	"github.com/iancoleman/strcase"
)

// Returns the segment of the path that is used to identify the resource.
// This is the same as the name of the struct that is used to represent the resource but
func (c *Client) GetSegment(v interface{}) string {
	// If v is a pointer, get the value it points to.
	if reflect.TypeOf(v).Kind() == reflect.Ptr {
		v = reflect.ValueOf(v).Elem().Interface()
	}

	// If v is a slice, get the type of the first element.
	if reflect.TypeOf(v).Kind() == reflect.Slice {
		if reflect.ValueOf(v).Len() == 0 {
			return "unknown"
		}
		v = reflect.ValueOf(v).Index(0).Interface()
	}

	t := reflect.TypeOf(v).Name()

	// If the type name contains "Collection", remove it.
	if t[len(t)-10:] == "Collection" {
		t = t[:len(t)-10]
	}

	return strcase.ToKebab(t)
}
