type Native struct {
	Request string `json:"request" binding:"required"`
	Ver     string `json:"ver,omitempty"`
	API     []int  `json:"api,omitempty"`
	BAttr   []int  `json:"battr,omitempty"`
	Ext     interface{} `json:"ext,omitempty"`
}