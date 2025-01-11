package gortb

import (
	"testing"
	"errors"
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
				Imp:  []Imp{{ID: "1"}},
				Site: &Site{},
			},
			wantErr: nil,
		},
		{
			name: "valid request with app",
			request: &BidRequest{
				ID:  "test-id",
				Imp: []Imp{{ID: "1"}},
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
		{
			name: "invalid test value",
			request: &BidRequest{
				ID:   "test-id",
				Imp:  []Imp{{ID: "1"}},
				Site: &Site{},
				Test: 2,
			},
			wantErr: ErrInvalidTest,
		},
		{
			name: "invalid auction type",
			request: &BidRequest{
				ID:   "test-id",
				Imp:  []Imp{{ID: "1"}},
				Site: &Site{},
				At:   3,
			},
			wantErr: ErrInvalidAt,
		},
		{
			name: "invalid allimps",
			request: &BidRequest{
				ID:      "test-id",
				Imp:     []Imp{{ID: "1"}},
				Site:    &Site{},
				AllImps: 2,
			},
			wantErr: ErrInvalidAllImps,
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
				t.Errorf("BidRequest.Validate() unexpected error = %v", err)
			}
		})
	}
}
