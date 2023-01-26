package model

type EntityCollection struct {
	Total        int64       `json:"total"`
	Aggregations interface{} `json:"aggregations"`
}
