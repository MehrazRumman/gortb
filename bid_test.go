package gortb

import (
	"encoding/json"
	"testing"
)

func TestBid_Validate(t *testing.T) {
	tests := []struct {
		name    string
		bid     *Bid
		wantErr error
	}{
		{
			name: "valid bid",
			bid: &Bid{
				ID:    "test_bid_id",
				ImpID: "test_imp_id",
				Price: 1.23,
			},
			wantErr: nil,
		},
		{
			name: "missing id",
			bid: &Bid{
				ImpID: "test_imp_id",
				Price: 1.23,
			},
			wantErr: ErrMissingBidID,
		},
		{
			name: "missing impression id",
			bid: &Bid{
				ID:    "test_bid_id",
				Price: 1.23,
			},
			wantErr: ErrMissingImpID,
		},
		{
			name: "invalid price",
			bid: &Bid{
				ID:    "test_bid_id",
				ImpID: "test_imp_id",
				Price: 0,
			},
			wantErr: ErrInvalidPrice,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.bid.Validate()
			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("Bid.Validate() error = nil, wantErr = %v", tt.wantErr)
					return
				}
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Bid.Validate() error = %v, wantErr = %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("Bid.Validate() error = %v, wantErr = nil", err)
			}
		})
	}
}

func TestBidWithJSON(t *testing.T) {
	jsonData := `{
		"id": "test_bid_id",
		"impid": "test_imp_id", 
		"price": 1.23,
		"adm": "<ad markup>",
		"adid": "ad123",
		"adomain": ["advertiser.com"],
		"iurl": "https://image.url",
		"cid": "campaign123",
		"crid": "creative123",
		"w": 300,
		"h": 250
	}`

	var bid Bid
	err := json.Unmarshal([]byte(jsonData), &bid)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	err = bid.Validate()
	if err != nil {
		t.Errorf("Bid.Validate() error = %v, wantErr = nil", err)
	}
}
