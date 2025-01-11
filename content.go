package gortb

type Content struct {
	ID                string   `json:"id,omitempty"`
	Episode           int      `json:"episode,omitempty"`
	Title             string   `json:"title,omitempty"`
	Series            string   `json:"series,omitempty"`
	Season            string   `json:"season,omitempty"`
	Artist            string   `json:"artist,omitempty"`
	Genre             string   `json:"genre,omitempty"`
	Album             string   `json:"album,omitempty"`
	ISRC              string   `json:"isrc,omitempty"`
	Producer          *Producer `json:"producer,omitempty"`
	URL               string   `json:"url,omitempty"`
	Cat               []string `json:"cat,omitempty"`
	Probq             int      `json:"probq,omitempty"`
	VideoQuality      int      `json:"videoquality,omitempty"` // Deprecated
	Context           int      `json:"context,omitempty"`
	ContentRating     string   `json:"contentrating,omitempty"`
	UserRating        string   `json:"userrating,omitempty"`
	QAGMediaRating    int      `json:"qagmediarating,omitempty"`
	Keywords          string   `json:"keywords,omitempty"`
	LiveStream        int      `json:"livestream,omitempty"`
	SourceRelationship int     `json:"sourcerelationship,omitempty"`
	Length            int      `json:"len,omitempty"`
	Language          string   `json:"language,omitempty"`
	Embeddable        int      `json:"embeddable,omitempty"`
	Data              []Data   `json:"data,omitempty"`
	Ext               interface{} `json:"ext,omitempty"`
}