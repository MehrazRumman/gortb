package gortb

import (
	"testing"
	"encoding/json"
)

func TestAudio(t *testing.T) {
	tests := []struct {
		name    string
		audio   *Audio
		wantErr error
	}{
		{
			name: "valid audio",
			audio: &Audio{
				MIMEs:      []string{"audio/mp3"},
				StartDelay: 0,
				Sequence:   1,
				Feed:       1,
				Stitched:   0,
				NVol:       50,
			},
			wantErr: nil,
		},
		{
			name: "missing MIMEs",
			audio: &Audio{
				StartDelay: 0,
			},
			wantErr: ErrMissingMIMEs,
		},
		{
			name: "invalid start delay",
			audio: &Audio{
				MIMEs:      []string{"audio/mp3"},
				StartDelay: -3,
			},
			wantErr: ErrInvalidStartDelay,
		},
		{
			name: "invalid sequence",
			audio: &Audio{
				MIMEs:    []string{"audio/mp3"},
				Sequence: 0,
			},
			wantErr: ErrInvalidSequence,
		},
		{
			name: "invalid feed",
			audio: &Audio{
				MIMEs: []string{"audio/mp3"},
				Feed:  4,
			},
			wantErr: ErrInvalidFeed,
		},
		{
			name: "invalid stitched",
			audio: &Audio{
				MIMEs:    []string{"audio/mp3"},
				Stitched: 2,
			},
			wantErr: ErrInvalidStitched,
		},
		{
			name: "invalid nvol",
			audio: &Audio{
				MIMEs: []string{"audio/mp3"},
				NVol:  101,
			},
			wantErr: ErrInvalidNVol,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.audio.Validate()
			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("Audio.Validate() error = nil, wantErr = %v", tt.wantErr)
					return
				}
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Audio.Validate() error = %v, wantErr = %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("Audio.Validate() error = %v, wantErr = nil", err)
			}
		})
	}
}

func TestAudioWithJSON(t *testing.T) {
	jsonData := `{
		"mimes": ["audio/mp3"],
		"minduration": 5,
		"maxduration": 30,
		"startdelay": 0,
		"sequence": 1,
		"feed": 1,
		"stitched": 0,
		"nvol": 50
	}`

	var audio Audio
	err := json.Unmarshal([]byte(jsonData), &audio)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	err = audio.Validate()
	if err != nil {
		t.Errorf("Audio.Validate() error = %v, wantErr = nil", err)
	}
}
