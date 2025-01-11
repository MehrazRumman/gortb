type Metric struct {
	Type   string  `json:"type" binding:"required"`
	Value  float64 `json:"value" binding:"required"`
	Vendor string  `json:"vendor,omitempty"`
	Ext    interface{} `json:"ext,omitempty"`
}