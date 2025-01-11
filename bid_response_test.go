package gortb

import (
	"testing"
)

func TestBidResponse(t *testing.T) {
	tests := []struct {
		name    string
		resp    *BidResponse
		wantErr error
	}{
		{
			name: "valid response",
			resp: &BidResponse{
				ID: "test-id",
				SeatBids: []SeatBid{
					{
						Bids: []Bid{
							{
								ID:    "1",
								ImpID: "1",
								Price: 1.0,
							},
						},
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "missing id",
			resp: &BidResponse{
				SeatBids: []SeatBid{
					{
						Bids: []Bid{
							{
								ID:    "1",
								ImpID: "1",
								Price: 1.0,
							},
						},
					},
				},
			},
			wantErr: ErrMissingResponseID,
		},
		{
			name: "missing seatbids",
			resp: &BidResponse{
				ID: "test-id",
			},
			wantErr: ErrMissingSeatBids,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.resp.Validate()
			if err != tt.wantErr {
				t.Errorf("BidResponse.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
