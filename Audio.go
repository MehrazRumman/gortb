package gortb

import (
	"errors"
)

// Validation errors for Audio
var (
	ErrMissingMIMEs = errors.New("audio missing required MIMEs")
	ErrInvalidStartDelay = errors.New("audio start delay must be >= -2")
	ErrInvalidSequence = errors.New("audio sequence must be >= 1")
	ErrInvalidFeed = errors.New("audio feed must be 1, 2 or 3")
	ErrInvalidStitched = errors.New("audio stitched must be 0 or 1")
	ErrInvalidNVol = errors.New("audio nvol must be between 0 and 100")
)

type Audio struct {
	MIMEs        []string `json:"mimes" binding:"required"`
	MinDuration  int      `json:"minduration,omitempty"`
	MaxDuration  int      `json:"maxduration,omitempty"`
	Protocols    []int    `json:"protocols,omitempty"` 
	StartDelay   int      `json:"startdelay,omitempty"` 
	Sequence     int      `json:"sequence,omitempty"`
	BAttr        []int    `json:"battr,omitempty"`
	MaxExtended  int      `json:"maxextended,omitempty"`
	MinBitrate   int      `json:"minbitrate,omitempty"`
	MaxBitrate   int      `json:"maxbitrate,omitempty"`
	Delivery     []int    `json:"delivery,omitempty"`
	CompanionAd  []Banner `json:"companionad,omitempty"`
	API          []int    `json:"api,omitempty"`
	CompanionType []int   `json:"companiontype,omitempty"`
	MaxSeq       int      `json:"maxseq,omitempty"`
	Feed         int      `json:"feed,omitempty"`
	Stitched     int      `json:"stitched,omitempty"`
	NVol         int      `json:"nvol,omitempty"`
	Ext          interface{} `json:"ext,omitempty"`
}

// Validate performs validation on the Audio object according to OpenRTB 2.5 rules
func (a *Audio) Validate() error {
	// Check required fields
	if len(a.MIMEs) == 0 {
		return ErrMissingMIMEs
	}

	// Validate start delay if present
	// -2 = End-card
	// -1 = Mid-roll
	// 0 = Pre-roll
	// >0 = Generic delay in seconds
	if a.StartDelay < -2 {
		return ErrInvalidStartDelay
	}

	// Validate sequence if present
	if a.Sequence != 0 && a.Sequence < 1 {
		return ErrInvalidSequence
	}

	// Validate feed type if present
	// 1 = Music Service
	// 2 = FM/AM Broadcast
	// 3 = Podcast
	if a.Feed != 0 && (a.Feed < 1 || a.Feed > 3) {
		return ErrInvalidFeed
	}

	// Validate stitched field if present
	if a.Stitched != 0 && a.Stitched != 1 {
		return ErrInvalidStitched
	}

	// Validate volume normalization if present
	if a.NVol != 0 && (a.NVol < 0 || a.NVol > 100) {
		return ErrInvalidNVol
	}

	return nil
}