package gortb

import (
	"errors"
)

var (
	ErrMissingMIMEsVideo = errors.New("video mimes are required")
	ErrMissingDimensions = errors.New("video width and height are required")
)

type Video struct {
	MIMEs          []string    `json:"mimes" binding:"required"`
	MinDuration    int         `json:"minduration,omitempty"`
	MaxDuration    int         `json:"maxduration,omitempty"`
	Protocols      []int       `json:"protocols,omitempty"`
	Protocol       int         `json:"protocol,omitempty"` // Deprecated
	W              int         `json:"w,omitempty" binding:"required"`
	H              int         `json:"h,omitempty" binding:"required"`
	StartDelay     int         `json:"startdelay,omitempty"`
	Placement      int         `json:"placement,omitempty"`
	Linearity      int         `json:"linearity,omitempty"`
	Skip           int         `json:"skip,omitempty"`
	SkipMin        int         `json:"skipmin,omitempty" default:"0"`
	SkipAfter      int         `json:"skipafter,omitempty" default:"0"`
	Sequence       int         `json:"sequence,omitempty"`
	BAttr          []int       `json:"battr,omitempty"`
	MaxExtended    int         `json:"maxextended,omitempty"`
	MinBitrate     int         `json:"minbitrate,omitempty"`
	MaxBitrate     int         `json:"maxbitrate,omitempty"`
	BoxingAllowed  int         `json:"boxingallowed,omitempty" default:"1"`
	PlaybackMethod []int       `json:"playbackmethod,omitempty"`
	PlayBackend    int         `json:"playbackend,omitempty"`
	Delivery       []int       `json:"delivery,omitempty"`
	Pos            int         `json:"pos,omitempty"`
	CompanionAd    []Banner    `json:"companionad,omitempty"`
	API            []int       `json:"api,omitempty"`
	CompanionType  []int       `json:"companiontype,omitempty"`
	Ext            interface{} `json:"ext,omitempty"`
}

// Validate checks if the Video object meets OpenRTB requirements
func (v *Video) Validate() error {
	if len(v.MIMEs) == 0 {
		return ErrMissingMIMEsVideo
	}

	if v.W == 0 || v.H == 0 {
		return ErrMissingDimensions
	}

	return nil
}
