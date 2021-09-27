package qq

type ItemFileType struct {
	Key      string
	TypeName string
	Br       int
	Pre      string
	Ext      string
}

func GetFileTypeList() []ItemFileType {
	ret := []ItemFileType{
		{
			Key:      "Size320Mp3",
			TypeName: "size_320mp3",
			Br:       320,
			Pre:      "M800",
			Ext:      "mp3",
		},
		{
			Key:      "Size192Aac",
			TypeName: "size_192aac",
			Br:       192,
			Pre:      "C600",
			Ext:      "m4a",
		},
		{
			Key:      "Size128Mp3",
			TypeName: "size_128mp3",
			Br:       128,
			Pre:      "M500",
			Ext:      "mp3",
		},
		{
			Key:      "Size96Aac",
			TypeName: "size_96aac",
			Br:       96,
			Pre:      "C400",
			Ext:      "m4a",
		},
		{
			Key:      "Size48Aac",
			TypeName: "size_48aac",
			Br:       48,
			Pre:      "C200",
			Ext:      "m4a",
		},
		{
			Key:      "Size24Aac",
			TypeName: "size_24aac",
			Br:       24,
			Pre:      "C100",
			Ext:      "m4a",
		},
	}

	return ret
}

type Album struct {
	ID             int64   `json:"id"`
	Mid            string  `json:"mid"`
	Name           string  `json:"name"`
	Pmid           string  `json:"pmid"`
	Subtitle       *string `json:"subtitle,omitempty"`
	TimePublic     *string `json:"time_public,omitempty"`
	Title          string  `json:"title"`
	TitleHighlight string  `json:"title_highlight"`
	Type           *int64  `json:"type,omitempty"`
	Uin            *int64  `json:"uin,omitempty"`
}

type File struct {
	B30S          int64  `json:"b_30s"`
	E30S          int64  `json:"e_30s"`
	HiresBitdepth int64  `json:"hires_bitdepth"`
	HiresSample   int64  `json:"hires_sample"`
	MediaMid      string `json:"media_mid"`
	Size128       int64  `json:"size_128"`
	Size128Mp3    int64  `json:"size_128mp3"`
	Size192AAC    int64  `json:"size_192aac"`
	Size192Ogg    int64  `json:"size_192ogg"`
	Size24AAC     int64  `json:"size_24aac"`
	Size320       int64  `json:"size_320"`
	Size320Mp3    int64  `json:"size_320mp3"`
	Size48AAC     int64  `json:"size_48aac"`
	Size96AAC     int64  `json:"size_96aac"`
	Size96Ogg     int64  `json:"size_96ogg"`
	SizeAAC       int64  `json:"size_aac"`
	SizeApe       int64  `json:"size_ape"`
	SizeDts       int64  `json:"size_dts"`
	SizeFLAC      int64  `json:"size_flac"`
	SizeHires     int64  `json:"size_hires"`
	SizeOgg       int64  `json:"size_ogg"`
	SizeTry       int64  `json:"size_try"`
	StrMediaMid   string `json:"strMediaMid"`
	TryBegin      int64  `json:"try_begin"`
	TryEnd        int64  `json:"try_end"`
	URL           string `json:"url"`
}

type Ksong struct {
	ID  int64  `json:"id"`
	Mid string `json:"mid"`
}

type Mv struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Title string `json:"title"`
	Vid   string `json:"vid"`
	VT    int64  `json:"vt"`
}

type Pay struct {
	PayDown    int64 `json:"pay_down"`
	PayMonth   int64 `json:"pay_month"`
	PayPlay    int64 `json:"pay_play"`
	PayStatus  int64 `json:"pay_status"`
	PriceAlbum int64 `json:"price_album"`
	PriceTrack int64 `json:"price_track"`
	TimeFree   int64 `json:"time_free"`
}

type Volume struct {
	Gain float64 `json:"gain"`
	Lra  float64 `json:"lra"`
	Peak float64 `json:"peak"`
}
