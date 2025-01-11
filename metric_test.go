package gortb

import (
	"encoding/json"
	"testing"
)

func TestMetric(t *testing.T) {
	tests := []struct {
		name    string
		metric  *Metric
		wantErr error
	}{
		{
			name: "valid metric",
			metric: &Metric{
				Type:   "viewability",
				Value:  0.85,
				Vendor: "vendor1",
			},
			wantErr: nil,
		},
		{
			name: "missing type",
			metric: &Metric{
				Value:  0.85,
				Vendor: "vendor1",
			},
			wantErr: ErrMissingType,
		},
		{
			name: "missing value",
			metric: &Metric{
				Type:   "viewability",
				Vendor: "vendor1",
			},
			wantErr: ErrMissingValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.metric.Validate()
			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("Metric.Validate() error = nil, wantErr = %v", tt.wantErr)
					return
				}
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Metric.Validate() error = %v, wantErr = %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("Metric.Validate() error = %v, wantErr = nil", err)
			}
		})
	}
}

func TestMetricWithJSON(t *testing.T) {
	jsonData := `{
		"type": "viewability",
		"value": 0.85,
		"vendor": "vendor1"
	}`

	var metric Metric
	err := json.Unmarshal([]byte(jsonData), &metric)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	err = metric.Validate()
	if err != nil {
		t.Errorf("Metric.Validate() error = %v, wantErr = nil", err)
	}
}
