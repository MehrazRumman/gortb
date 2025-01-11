package gortb

import (
	"testing"
	"encoding/json"
)

func TestDeal(t *testing.T) {
	tests := []struct {
		name    string
		deal    *Deal
		wantErr error
	}{
		{
			name: "valid deal",
			deal: &Deal{
				ID:          "deal-123",
				BidFloor:    1.50,
				BidFloorCur: "USD",
				At:          2,
				WSeat:       []string{"seat1", "seat2"},
				WAdmin:      []string{"admin1"},
			},
			wantErr: nil,
		},
		{
			name: "missing ID",
			deal: &Deal{
				BidFloor:    1.50,
				BidFloorCur: "USD",
			},
			wantErr: ErrMissingDealID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.deal.Validate()
			if err != tt.wantErr {
				t.Errorf("Deal.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDealWithJSON(t *testing.T) {
	jsonData := `{
		"id": "deal-123",
		"bidfloor": 1.50,
		"bidfloorcur": "USD",
		"at": 2,
		"wseat": ["seat1", "seat2"],
		"wadmin": ["admin1"]
	}`

	var deal Deal
	err := json.Unmarshal([]byte(jsonData), &deal)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	err = deal.Validate()
	if err != nil {
		t.Errorf("Deal.Validate() error = %v, wantErr = nil", err)
	}
}
