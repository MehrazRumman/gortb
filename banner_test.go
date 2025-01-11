package gortb

import (
	"encoding/json"
	"testing"
)

func TestBanner_Validate(t *testing.T) {
	tests := []struct {
		name    string
		banner  *Banner
		wantErr error
	}{
		{
			name: "valid banner with w/h",
			banner: &Banner{
				W: 300,
				H: 250,
			},
			wantErr: nil,
		},
		{
			name: "valid banner with format",
			banner: &Banner{
				Format: []Format{
					{W: 300, H: 250},
				},
			},
			wantErr: nil,
		},
		{
			name: "invalid dimensions",
			banner: &Banner{
				W: 0,
				H: 0,
			},
			wantErr: ErrInvalidBannerDimensions,
		},
		{
			name: "invalid topframe",
			banner: &Banner{
				W:        300,
				H:        250,
				TopFrame: 2,
			},
			wantErr: ErrInvalidTopFrame,
		},
		{
			name: "invalid vcm",
			banner: &Banner{
				W:   300,
				H:   250,
				VCM: 2,
			},
			wantErr: ErrInvalidVCM,
		},
		{
			name: "invalid expdir low",
			banner: &Banner{
				W:      300,
				H:      250,
				ExpDir: 0,
			},
			wantErr: nil,
		},
		{
			name: "invalid expdir high",
			banner: &Banner{
				W:      300,
				H:      250,
				ExpDir: 7,
			},
			wantErr: ErrInvalidExpDir,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.banner.Validate()
			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("Banner.Validate() error = nil, wantErr = %v", tt.wantErr)
					return
				}
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Banner.Validate() error = %v, wantErr = %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("Banner.Validate() error = %v, wantErr = nil", err)
			}
		})
	}
}

func TestBannerWithJSON(t *testing.T) {
	jsonData := `{
		"w": 300,
		"h": 250,
		"id": "test_banner",
		"topframe": 1,
		"vcm": 0,
		"expdir": 1
	}`

	var banner Banner
	err := json.Unmarshal([]byte(jsonData), &banner)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	err = banner.Validate()
	if err != nil {
		t.Errorf("Banner.Validate() error = %v, wantErr = nil", err)
	}
}
