type Format struct {
	W       int `json:"w"`
	H       int `json:"h"`
	WRatio  int `json:"wratio,omitempty"`
	HRatio  int `json:"hratio,omitempty"`
	WMin    int `json:"wmin,omitempty"`
	Ext     interface{} `json:"ext,omitempty"`
}