package gortb

import (
	"testing"
	"encoding/json"
)

func TestImp(t *testing.T) {
	tests := []struct {
		name    string
		imp     *Imp
		wantErr error
	}{
		{
			name: "valid banner impression",
			imp: &Imp{
				ID:     "test-imp-1",
				Banner: &Banner{},
			},
			wantErr: nil,
		},
		{
			name: "valid video impression",
			imp: &Imp{
				ID:    "test-imp-2",
				Video: &Video{},
			},
			wantErr: nil,
		},
		{
			name: "valid audio impression",
			imp: &Imp{
				ID:    "test-imp-3",
				Audio: &Audio{},
			},
			wantErr: nil,
		},
		{
			name: "valid native impression",
			imp: &Imp{
				ID:     "test-imp-4",
				Native: &Native{},
			},
			wantErr: nil,
		},
		{
			name: "missing ID",
			imp: &Imp{
				Banner: &Banner{},
			},
			wantErr: ErrMissingID,
		},
		{
			name: "invalid instl value",
			imp: &Imp{
				ID:    "test-imp-5",
				Instl: 2,
				Banner: &Banner{},
			},
			wantErr: ErrInvalidInstl,
		},
		{
			name: "invalid clickbrowser value",
			imp: &Imp{
				ID:           "test-imp-6",
				ClickBrowser: 2,
				Banner:      &Banner{},
			},
			wantErr: ErrInvalidClickBrowser,
		},
		{
			name: "invalid secure value",
			imp: &Imp{
				ID:     "test-imp-7",
				Secure: 2,
				Banner: &Banner{},
			},
			wantErr: ErrInvalidSecure,
		},
		{
			name: "negative bidfloor",
			imp: &Imp{
				ID:       "test-imp-8",
				BidFloor: -1.0,
				Banner:   &Banner{},
			},
			wantErr: ErrInvalidBidFloor,
		},
		{
			name: "missing bidfloorcur with non-zero bidfloor",
			imp: &Imp{
				ID:       "test-imp-9",
				BidFloor: 1.0,
				BidFloorCur: "",
				Banner:   &Banner{},
			},
			wantErr: ErrMissingBidFloorCur,
		},
		{
			name: "no media present",
			imp: &Imp{
				ID: "test-imp-10",
			},
			wantErr: ErrNoMediaPresent,
		},
		{
			name: "multiple media present",
			imp: &Imp{
				ID:     "test-imp-11",
				Banner: &Banner{},
				Video:  &Video{},
			},
			wantErr: ErrMultipleMedia,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.imp.Validate()
			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("Imp.Validate() error = nil, wantErr = %v", tt.wantErr)
					return
				}
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Imp.Validate() error = %v, wantErr = %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("Imp.Validate() error = %v, wantErr = nil", err)
			}
		})
	}
}

func TestImpWithJSON(t *testing.T) {
	jsonData := `{
		"id": "test-imp-1",
		"banner": {},
		"bidfloor": 1.0,
		"bidfloorcur": "USD"
	}`

	var imp Imp
	err := json.Unmarshal([]byte(jsonData), &imp)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	err = imp.Validate()
	if err != nil {
		t.Errorf("Imp.Validate() error = %v, wantErr = nil", err)
	}
}
