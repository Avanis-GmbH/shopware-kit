package com

import (
	"reflect"
	"strings"

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

	// If the type name contains "Collection", remove it.
	return strcase.ToKebab(strings.TrimSuffix(reflect.TypeOf(v).Name(), "Collection"))
}

func (c *Client) GetSegmentSnakeCase(v interface{}) string {
	// If the type name contains "Collection", remove it.
	return strcase.ToSnake(c.GetSegment(v))
}
