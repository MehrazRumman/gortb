package gortb

import "errors"

var (
	ErrMissingType  = errors.New("metric type is required")
	ErrMissingValue = errors.New("metric value is required")
)

type Metric struct {
	Type   string      `json:"type" binding:"required"`
	Value  float64     `json:"value" binding:"required"`
	Vendor string      `json:"vendor,omitempty"`
	Ext    interface{} `json:"ext,omitempty"`
}

// Validate performs validation on the Metric object according to OpenRTB 2.5 rules
func (m *Metric) Validate() error {
	// Check required fields
	if m.Type == "" {
		return ErrMissingType
	}

	// Value is required and implicitly validated by binding tag
	if m.Value == 0 {
		return ErrMissingValue
	}

	return nil
}
