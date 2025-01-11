package gortb

import (
	"errors"
)

// Validation errors for Banner
var (
	ErrInvalidBannerDimensions = errors.New("banner must have either w/h or format specified")
	ErrInvalidTopFrame = errors.New("banner topframe must be 0 or 1")
	ErrInvalidVCM = errors.New("banner vcm must be 0 or 1")
	ErrInvalidExpDir = errors.New("banner expdir must be between 1 and 6")
)

type Banner struct {
	Format       []Format `json:"format,omitempty"`
	W            int      `json:"w"`
	H            int      `json:"h"`
	WMax         int      `json:"wmax,omitempty"`    // Deprecated
	HMax         int      `json:"hmax,omitempty"`    // Deprecated
	WMin         int      `json:"wmin,omitempty"`    // Deprecated
	HMin         int      `json:"hmin,omitempty"`    // Deprecated
	BType        []int    `json:"btype,omitempty"`
	BAttr        []int    `json:"battr,omitempty"`
	Pos          int      `json:"pos,omitempty"`
	Mimes        []string `json:"mimes,omitempty"`
	TopFrame     int      `json:"topframe,omitempty"`
	ExpDir       int      `json:"expdir,omitempty"`
	API          []int    `json:"api,omitempty"`
	ID           string   `json:"id,omitempty"`
	VCM          int      `json:"vcm,omitempty"`
	Ext          interface{} `json:"ext,omitempty"`
}

// Validate performs validation on the Banner object according to OpenRTB 2.5 rules
func (b *Banner) Validate() error {
	// Check that either w/h or format is specified
	if b.W == 0 && b.H == 0 && len(b.Format) == 0 {
		return ErrInvalidBannerDimensions
	}

	// Validate topframe if present (0 = below, 1 = above)
	if b.TopFrame != 0 && b.TopFrame != 1 {
		return ErrInvalidTopFrame
	}

	// Validate vcm if present (0 = no, 1 = yes)
	if b.VCM != 0 && b.VCM != 1 {
		return ErrInvalidVCM
	}

	// Validate expdir if present (1-6 are valid values)
	if b.ExpDir != 0 && (b.ExpDir < 1 || b.ExpDir > 6) {
		return ErrInvalidExpDir
	}

	return nil
}