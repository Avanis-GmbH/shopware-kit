package com

// TotalCountMode indicates how the total count of the search result should be calculated.
type TotalCountMode uint

// SearchFilterType is the type of the filter to be applied when searching.
type SearchFilterType string

// SearchSortDirection is the direction of the sort to be applied when searching.
type SearchSortDirection string

const (
	TotalCountModeDefault  = 1
	TotalCountModeExact    = 2
	TotalCountModeNextPage = 3

	SearchFilterTypeEquals    = "equals"
	SearchFilterTypeEqualsAny = "equalsAny"
	SearchFilterTypeContains  = "contains"
	SearchFilterTypeRange     = "range"
	SearchFilterTypeNot       = "not"
	SearchFilterTypeMulti     = "multi"
	SearchFilterTypePrefix    = "prefix"
	SearchFilterTypeSuffix    = "suffix"

	SearchSortDirectionAscending  = "ASC"
	SearchSortDirectionDescending = "DESC"
)

// Criteria is the struct that sums up all the search criteria.
type Criteria struct {
	Includes       map[string][]string `json:"includes,omitempty"`
	Page           int64               `json:"page,omitempty"`
	Limit          int64               `json:"limit,omitempty"`
	IDs            []string            `json:"ids,omitempty"`
	Filter         []interface{}       `json:"filter,omitempty"`
	PostFilter     []interface{}       `json:"postFilter,omitempty"`
	Sort           []CriteriaSort      `json:"sort,omitempty"`
	Associations   map[string]Criteria `json:"associations,omitempty"`
	Term           string              `json:"term,omitempty"`
	TotalCountMode int                 `json:"totalCountMode,omitempty"`
	Query          []CriteriaQuery     `json:"query,omitempty"`
}

// CriteriaFilter is the struct that defines a filter to be applied when searching.
type CriteriaFilter struct {
	Type     string        `json:"type"`
	Operator string        `json:"operator,omitempty"`
	Queries  []interface{} `json:"queries,omitempty"`
}

// CriteriaFilterQuery is an extension of the CriteriaFilter struct that includes the field and value to be filtered.
type CriteriaFilterQuery struct {
	CriteriaFilter
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}

// CriteriaQuery is a struct that defines queries and a score to be applied when searching.
type CriteriaQuery struct {
	Score uint64              `json:"score"`
	Query CriteriaFilterQuery `json:"query"`
}

// CriteriaSort is the struct that defines a sort to be applied when searching.
type CriteriaSort struct {
	Direction      string `json:"order"`
	Field          string `json:"field"`
	NaturalSorting bool   `json:"naturalSorting"`
}
