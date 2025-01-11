package gortb

import (
	"errors"
)

// Validation errors for Audio
var (
	ErrMissingMIMEsAudio = errors.New("audio missing required MIMEs")
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
		return ErrMissingMIMEsAudio
	}

	return nil
}