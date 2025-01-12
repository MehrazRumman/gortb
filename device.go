package gortb

type Device struct {
	UA             string      `json:"ua,omitempty"`
	Geo            *Geo        `json:"geo,omitempty" `
	DNT            int         `json:"dnt,omitempty"`
	LMT            int         `json:"lmt,omitempty" `
	IP             string      `json:"ip,omitempty" `
	IPv6           string      `json:"ipv6,omitempty"`
	DeviceType     int         `json:"devicetype,omitempty"`
	Make           string      `json:"make,omitempty"`
	Model          string      `json:"model,omitempty"`
	OS             string      `json:"os,omitempty"`
	OSV            string      `json:"osv,omitempty"`
	HWV            string      `json:"hwv,omitempty"`
	H              int         `json:"h,omitempty"`
	W              int         `json:"w,omitempty"`
	PPI            int         `json:"ppi,omitempty"`
	PXRatio        float64     `json:"pxratio,omitempty"`
	JS             int         `json:"js,omitempty"`
	GeoFetch       int         `json:"geofetch,omitempty"`
	FlashVer       string      `json:"flashver,omitempty"`
	Language       string      `json:"language,omitempty"`
	Carrier        string      `json:"carrier,omitempty"`
	MCCMNC         string      `json:"mccmnc,omitempty"`
	ConnectionType int         `json:"connectiontype,omitempty"`
	IFA            string      `json:"ifa,omitempty"`
	DIDSHA1        string      `json:"didsha1,omitempty"`
	DIDMD5         string      `json:"didmd5,omitempty"`
	DPIDSHA1       string      `json:"dpidsha1,omitempty"`
	DPIDMD5        string      `json:"dpidmd5,omitempty"`
	MacSHA1        string      `json:"macsha1,omitempty"`
	MacMD5         string      `json:"macmd5,omitempty"`
	Ext            interface{} `json:"ext,omitempty"`
}
