package gortb

import (
	"testing"
	"encoding/json"
)

func TestApp(t *testing.T) {
	tests := []struct {
		name    string
		app     *App
		wantErr error
	}{
		{
			name: "valid app",
			app: &App{
				ID:            "test-app-1",
				PrivacyPolicy: 1,
				Paid:          0,
			},
			wantErr: nil,
		},
		{
			name: "missing id",
			app: &App{
				PrivacyPolicy: 1,
				Paid:          0,
			},
			wantErr: ErrMissingAppID,
		},
		{
			name: "invalid privacy policy",
			app: &App{
				ID:            "test-app-2",
				PrivacyPolicy: 2,
				Paid:          0,
			},
			wantErr: ErrInvalidPrivacyPolicy,
		},
		{
			name: "invalid paid",
			app: &App{
				ID:            "test-app-3",
				PrivacyPolicy: 1,
				Paid:          2,
			},
			wantErr: ErrInvalidPaid,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.app.Validate()
			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("App.Validate() error = nil, wantErr = %v", tt.wantErr)
					return
				}
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("App.Validate() error = %v, wantErr = %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("App.Validate() error = %v, wantErr = nil", err)
			}
		})
	}
}

func TestAppWithJSON(t *testing.T) {
	jsonData := `{
		"id": "test-app-1",
		"name": "Test App",
		"bundle": "com.test.app",
		"privacypolicy": 1,
		"paid": 0
	}`

	var app App
	err := json.Unmarshal([]byte(jsonData), &app)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	err = app.Validate()
	if err != nil {
		t.Errorf("App.Validate() error = %v, wantErr = nil", err)
	}
}
