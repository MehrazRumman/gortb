package gortb




type Imp struct {
	ID           string    `json:"id" binding:"required"`
	Metric       []Metric  `json:"metric,omitempty"`
	Banner       *Banner   `json:"banner,omitempty"`
	Video        *Video    `json:"video,omitempty"`
	Audio        *Audio    `json:"audio,omitempty"`
	Native       *Native   `json:"native,omitempty"`
	PMP          *PMP      `json:"pmp,omitempty"`
	DisplayManager string   `json:"displaymanager,omitempty"`
	DisplayManagerVer string `json:"displaymanagerver,omitempty"`
	Instl        int       `json:"instl,omitempty" default:"0"`
	TagID        string    `json:"tagid,omitempty"`
	BidFloor     float64   `json:"bidfloor,omitempty" default:"0"`
	BidFloorCur  string    `json:"bidfloorcur,omitempty" default:"USD"`
	ClickBrowser int       `json:"clickbrowser,omitempty"`
	Secure       int       `json:"secure,omitempty"`
	IFrameBuster []string  `json:"iframebuster,omitempty"`
	Exp          int       `json:"exp,omitempty"`
	Ext          interface{} `json:"ext,omitempty"`
}