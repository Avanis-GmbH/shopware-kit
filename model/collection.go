package model

type Collection interface {
	SetTotal(int64)
	GetTotal() int64
	SetAggregations(interface{})
	GetAggregations() interface{}
	SetData([]interface{})
	GetData() []interface{}
}

type EntityCollection struct {
	Total        int64       `json:"total"`
	Aggregations interface{} `json:"aggregations"`

	Data []interface{} `json:"data"`
}

func (c EntityCollection) SetTotal(total int64) {
	c.Total = total
}

func (c EntityCollection) GetTotal() int64 {
	return c.Total
}

func (c EntityCollection) SetAggregations(aggregations interface{}) {
	c.Aggregations = aggregations
}

func (c EntityCollection) GetAggregations() interface{} {
	return c.Aggregations
}

func (c EntityCollection) SetData(data []interface{}) {
	c.Data = data
}

func (c EntityCollection) GetData() []interface{} {
	return c.Data
}
