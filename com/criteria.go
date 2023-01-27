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

	SearchSortDirectionAscending  = "ASC"
	SearchSortDirectionDescending = "DESC"
)

// Criteria is the struct that sums up all the search criteria.
type Criteria struct {
	Includes       map[string][]string `json:"includes,omitempty"`
	Page           int64               `json:"page,omitempty"`
	Limit          int64               `json:"limit,omitempty"`
	IDs            []string            `json:"ids,omitempty"`
	Filter         []CriteriaFilter    `json:"filter,omitempty"`
	PostFilter     []CriteriaFilter    `json:"postFilter,omitempty"`
	Sort           []CriteriaSort      `json:"sort,omitempty"`
	Associations   map[string]Criteria `json:"associations,omitempty"`
	Term           string              `json:"term,omitempty"`
	TotalCountMode int                 `json:"totalCountMode,omitempty"`
}

// CriteriaFilter is the struct that defines a filter to be applied when searching.
type CriteriaFilter struct {
	Type  string      `json:"type"`
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}

// CriteriaSort is the struct that defines a sort to be applied when searching.
type CriteriaSort struct {
	Direction      string `json:"order"`
	Field          string `json:"field"`
	NaturalSorting bool   `json:"naturalSorting"`
}
