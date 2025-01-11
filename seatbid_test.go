package gortb

import (
	"encoding/json"
	"testing"
)

func TestSeatBid(t *testing.T) {
	tests := []struct {
		name    string
		seatbid *SeatBid
		wantErr error
	}{
		{
			name: "valid seatbid",
			seatbid: &SeatBid{
				Bids: []Bid{
					{
						ID:     "test-bid-1",
						ImpID:  "test-imp-1",
						Price:  1.23,
						AdID:   "test-ad-1",
						NURL:   "http://test.com/nurl",
						BURL:   "http://test.com/burl",
						LURL:   "http://test.com/lurl",
						AdM:    "<creative>test</creative>",
						CID:    "test-campaign-1",
					},
				},
				Seat:  "test-seat",
				Group: 0,
			},
			wantErr: nil,
		},
		{
			name: "empty bids",
			seatbid: &SeatBid{
				Bids:  []Bid{},
				Seat:  "test-seat",
				Group: 0,
			},
			wantErr: ErrInvalidBids,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.seatbid.Validate()
			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("SeatBid.Validate() error = nil, wantErr = %v", tt.wantErr)
					return
				}
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("SeatBid.Validate() error = %v, wantErr = %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("SeatBid.Validate() error = %v, wantErr = nil", err)
			}
		})
	}
}

func TestSeatBidWithJSON(t *testing.T) {
	jsonData := `{
		"bid": [{
			"id": "test-bid-1",
			"impid": "test-imp-1",
			"price": 1.23,
			"adid": "test-ad-1",
			"nurl": "http://test.com/nurl",
			"burl": "http://test.com/burl", 
			"lurl": "http://test.com/lurl",
			"adm": "<creative>test</creative>",
			"cid": "test-campaign-1",
			"crid": "test-creative-1",
			"apikey": "test-api-key"
		}],
		"seat": "test-seat",
		"group": 0
	}`

	var seatbid SeatBid
	err := json.Unmarshal([]byte(jsonData), &seatbid)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	err = seatbid.Validate()
	if err != nil {
		t.Errorf("SeatBid.Validate() error = %v, wantErr = nil", err)
	}
}
