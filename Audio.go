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