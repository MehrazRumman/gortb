package gortb

import(
	"errors"
)

var (
	ErrMissingRequest = errors.New("native request field is required")
)

type Native struct {
	Request string `json:"request" binding:"required"`
	Ver     string `json:"ver,omitempty"`
	API     []int  `json:"api,omitempty"`
	BAttr   []int  `json:"battr,omitempty"`
	Ext     interface{} `json:"ext,omitempty"`
}

// Validate performs validation on the Native object according to OpenRTB 2.5 rules
func (n *Native) Validate() error {
	// Check required request field
	if n.Request == "" {
		return ErrMissingRequest
	}
	return nil
}