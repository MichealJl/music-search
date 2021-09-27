package qq

type GetPlayInfoRsp struct {
	Code      int64         `json:"code"`
	Data      []Datum       `json:"data"`
	URL       URL           `json:"url"`
	Url1      Url1          `json:"url1"`
	ExtraData []interface{} `json:"extra_data"`
	Joox      int64         `json:"joox"`
	JooxLogin int64         `json:"joox_login"`
	Msgid     int64         `json:"msgid"`
}

type Datum struct {
	Action      map[string]int64 `json:"action"`
	Aid         int64            `json:"aid"`
	Album       Album            `json:"album"`
	BPM         int64            `json:"bpm"`
	DataType    int64            `json:"data_type"`
	Es          string           `json:"es"`
	File        File             `json:"file"`
	Fnote       int64            `json:"fnote"`
	Genre       int64            `json:"genre"`
	ID          int64            `json:"id"`
	IndexAlbum  int64            `json:"index_album"`
	IndexCD     int64            `json:"index_cd"`
	Interval    int64            `json:"interval"`
	Isonly      int64            `json:"isonly"`
	Ksong       Ksong            `json:"ksong"`
	Label       string           `json:"label"`
	Language    int64            `json:"language"`
	Mid         string           `json:"mid"`
	ModifyStamp int64            `json:"modify_stamp"`
	Mv          Mv               `json:"mv"`
	Name        string           `json:"name"`
	Ov          int64            `json:"ov"`
	Pay         Pay              `json:"pay"`
	Sa          int64            `json:"sa"`
	Singer      []Singer         `json:"singer"`
	Status      int64            `json:"status"`
	Subtitle    string           `json:"subtitle"`
	Tid         int64            `json:"tid"`
	TimePublic  string           `json:"time_public"`
	Title       string           `json:"title"`
	Trace       string           `json:"trace"`
	Type        int64            `json:"type"`
	URL         string           `json:"url"`
	Version     int64            `json:"version"`
	Volume      Volume           `json:"volume"`
}

type Singer struct {
	ID    int64  `json:"id"`
	Mid   string `json:"mid"`
	Name  string `json:"name"`
	Pmid  string `json:"pmid"`
	Title string `json:"title"`
	Type  int64  `json:"type"`
	Uin   int64  `json:"uin"`
}

type URL struct {
	The320657613 string `json:"320657613"`
}

type Url1 struct {
}
