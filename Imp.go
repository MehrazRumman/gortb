package gortb

import "errors"

var (
	ErrInvalidInstl        = errors.New("invalid instl value, must be 0 or 1")
	ErrInvalidClickBrowser = errors.New("invalid clickbrowser value, must be 0 or 1")
	ErrInvalidSecure       = errors.New("invalid secure value, must be 0 or 1")
	ErrNoMediaPresent      = errors.New("impression must have at least one of: banner, video, audio, or native")
	ErrMultipleMedia       = errors.New("impression cannot have more than one of: banner, video, audio, or native")
	ErrInvalidBidFloor     = errors.New("bidfloor must be >= 0")
	ErrMissingBidFloorCur  = errors.New("bidfloorcur is required when bidfloor > 0")
)

type Imp struct {
	ID                string      `json:"id" binding:"required"`
	Metric            []Metric    `json:"metric,omitempty"`
	Banner            *Banner     `json:"banner,omitempty"`
	Video             *Video      `json:"video,omitempty"`
	Audio             *Audio      `json:"audio,omitempty"`
	Native            *Native     `json:"native,omitempty"`
	PMP               *PMP        `json:"pmp,omitempty"`
	DisplayManager    string      `json:"displaymanager,omitempty"`
	DisplayManagerVer string      `json:"displaymanagerver,omitempty"`
	Instl             int         `json:"instl,omitempty" default:"0"`
	TagID             string      `json:"tagid,omitempty"`
	BidFloor          float64     `json:"bidfloor,omitempty" default:"0"`
	BidFloorCur       string      `json:"bidfloorcur,omitempty" default:"USD"`
	ClickBrowser      int         `json:"clickbrowser,omitempty"`
	Secure            int         `json:"secure,omitempty"`
	IFrameBuster      []string    `json:"iframebuster,omitempty"`
	Exp               int         `json:"exp,omitempty"`
	Ext               interface{} `json:"ext,omitempty"`
}

// Validate performs validation on the Imp object according to OpenRTB 2.5 rules
func (imp *Imp) Validate() error {
	// Check required fields
	if imp.ID == "" {
		return ErrMissingID
	}

	// Validate instl field
	if imp.Instl != 0 && imp.Instl != 1 {
		return ErrInvalidInstl
	}

	// Validate clickbrowser field
	if imp.ClickBrowser != 0 && imp.ClickBrowser != 1 {
		return ErrInvalidClickBrowser
	}

	// Validate secure field
	if imp.Secure != 0 && imp.Secure != 1 {
		return ErrInvalidSecure
	}

	// Validate bidfloor and bidfloorcur
	if imp.BidFloor < 0 {
		return ErrInvalidBidFloor
	}
	if imp.BidFloor > 0 && imp.BidFloorCur == "" {
		return ErrMissingBidFloorCur
	}

	// Check that at least one media type is present
	mediaCount := 0
	if imp.Banner != nil {
		mediaCount++
	}
	if imp.Video != nil {
		mediaCount++
	}
	if imp.Audio != nil {
		mediaCount++
	}
	if imp.Native != nil {
		mediaCount++
	}

	if mediaCount == 0 {
		return ErrNoMediaPresent
	}
	if mediaCount > 1 {
		return ErrMultipleMedia
	}

	return nil
}
