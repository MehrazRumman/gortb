package gortb

type Geo struct {
	Lat           float64     `json:"lat,omitempty"`
	Lon           float64     `json:"lon,omitempty"`
	Type          int         `json:"type,omitempty"`
	Accuracy      int         `json:"accuracy,omitempty"`
	LastFix       int         `json:"lastfix,omitempty"`
	IPService     int         `json:"ipservice,omitempty"`
	Country       string      `json:"country,omitempty"`
	Region        string      `json:"region,omitempty"`
	RegionFIPS104 string      `json:"regionfips104,omitempty"`
	Metro         string      `json:"metro,omitempty"`
	City          string      `json:"city,omitempty"`
	Zip           string      `json:"zip,omitempty"`
	UTCOffset     int         `json:"utcoffset,omitempty"`
	Ext           interface{} `json:"ext,omitempty"`
}