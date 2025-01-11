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