package gortb

import (
	"testing"
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

func TestDealDefaults(t *testing.T) {
	deal := &Deal{
		ID: "deal-123",
	}

	// Test default values
	if deal.BidFloorCur != "USD" {
		t.Errorf("Deal default BidFloorCur = %v, want USD", deal.BidFloorCur)
	}

	if deal.BidFloor != 0 {
		t.Errorf("Deal default BidFloor = %v, want 0", deal.BidFloor)
	}
}
