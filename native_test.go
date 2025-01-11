package gortb

import (
	"testing"
	"encoding/json"
)

func TestNative(t *testing.T) {
	tests := []struct {
		name    string
		native  *Native
		wantErr error
	}{
		{
			name: "valid native",
			native: &Native{
				Request: "native_request_string",
				Ver:     "1.0",
				API:     []int{1, 2},
				BAttr:   []int{1, 2, 3},
			},
			wantErr: nil,
		},
		{
			name: "missing request",
			native: &Native{
				Ver:   "1.0",
				API:   []int{1, 2},
				BAttr: []int{1, 2, 3},
			},
			wantErr: ErrMissingRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.native.Validate()
			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("Native.Validate() error = nil, wantErr = %v", tt.wantErr)
					return
				}
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Native.Validate() error = %v, wantErr = %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("Native.Validate() error = %v, wantErr = nil", err)
			}
		})
	}
}

func TestNativeWithJSON(t *testing.T) {
	jsonData := `{
		"request": "native_request_string",
		"ver": "1.0",
		"api": [1, 2],
		"battr": [1, 2, 3],
		"ext": {
			"custom_field": "value"
		}
	}`

	var native Native
	err := json.Unmarshal([]byte(jsonData), &native)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	err = native.Validate()
	if err != nil {
		t.Errorf("Native.Validate() error = %v, wantErr = nil", err)
	}
}
