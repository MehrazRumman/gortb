package gortb

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestBidRequest(t *testing.T) {
	tests := []struct {
		name    string
		request *BidRequest
		wantErr error
	}{
		{
			name: "valid request with site",
			request: &BidRequest{
				ID:   "test-id",
				Imp:  []Imp{{ID: "1", Banner: &Banner{}}},
				Site: &Site{},
			},
			wantErr: nil,
		},
		{
			name: "valid request with app",
			request: &BidRequest{
				ID:  "test-id",
				Imp: []Imp{{ID: "1", Banner: &Banner{}}},
				App: &App{},
			},
			wantErr: nil,
		},
		{
			name: "missing id",
			request: &BidRequest{
				Imp:  []Imp{{ID: "1"}},
				Site: &Site{},
			},
			wantErr: ErrMissingID,
		},
		{
			name: "missing impressions",
			request: &BidRequest{
				ID:   "test-id",
				Site: &Site{},
			},
			wantErr: ErrMissingImp,
		},
		{
			name: "both site and app",
			request: &BidRequest{
				ID:   "test-id",
				Imp:  []Imp{{ID: "1"}},
				Site: &Site{},
				App:  &App{},
			},
			wantErr: ErrSiteAndApp,
		},
		{
			name: "neither site nor app",
			request: &BidRequest{
				ID:  "test-id",
				Imp: []Imp{{ID: "1"}},
			},
			wantErr: errors.New("bid request must contain either site or app"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("BidRequest.Validate() error = nil, wantErr = %v", tt.wantErr)
					return
				}
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("BidRequest.Validate() error = %v, wantErr = %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("BidRequest.Validate() error = %v, wantErr = nil", err)
			}
		})
	}
}

func TestBidRequestWithJSON(t *testing.T) {
	jsonData := `{
		"id": "test-id",
		"imp": [
			{
				"id": "1",
				"banner": {}
			}
		],
		"site": {}
	}`

	var request BidRequest
	err := json.Unmarshal([]byte(jsonData), &request)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Validate the BidRequest
	err = request.Validate()
	if err != nil {
		t.Errorf("BidRequest.Validate() error = %v, wantErr = nil", err)
	}
}
