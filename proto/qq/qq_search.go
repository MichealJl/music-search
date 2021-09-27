package qq

// QQSearchRsp 搜索音乐返回结果
type QQSearchRsp struct {
	Code    int64  `json:"code"`
	Data    QQData `json:"data"`
	Message string `json:"message"`
	Notice  string `json:"notice"`
	Subcode int64  `json:"subcode"`
	Time    int64  `json:"time"`
	Tips    string `json:"tips"`
}

type QQData struct {
	Keyword   string        `json:"keyword"`
	Priority  int64         `json:"priority"`
	Qc        []interface{} `json:"qc"`
	Semantic  Semantic      `json:"semantic"`
	Song      Semantic      `json:"song"`
	Tab       int64         `json:"tab"`
	Taglist   []interface{} `json:"taglist"`
	Totaltime int64         `json:"totaltime"`
	Zhida     Zhida         `json:"zhida"`
}

type Semantic struct {
	Curnum   int64        `json:"curnum"`
	Curpage  int64        `json:"curpage"`
	List     []QQDataItem `json:"list"`
	Totalnum int64        `json:"totalnum"`
}

type QQDataItem struct {
	Act          int64            `json:"act"`
	Action  map[string]int64 `json:"action"`
	Album   Album            `json:"album"`
	BPM     int64            `json:"bpm"`
	Content      string           `json:"content"`
	Desc         string           `json:"desc"`
	DescHilight  string           `json:"desc_hilight"`
	Docid        string           `json:"docid"`
	Eq           int64            `json:"eq"`
	Es      string           `json:"es"`
	File    File             `json:"file"`
	Fnote   int64            `json:"fnote"`
	Genre        int64            `json:"genre"`
	Grp          []interface{}    `json:"grp"`
	Href3        string           `json:"href3"`
	ID           int64            `json:"id"`
	IndexAlbum   int64            `json:"index_album"`
	IndexCD      int64            `json:"index_cd"`
	Interval     int64            `json:"interval"`
	Isonly  int64            `json:"isonly"`
	Ksong   Ksong            `json:"ksong"`
	Label   string           `json:"label"`
	Language     int64            `json:"language"`
	Lyric        string           `json:"lyric"`
	LyricHilight string           `json:"lyric_hilight"`
	Mid     string           `json:"mid"`
	Mv      Mv               `json:"mv"`
	Name    string           `json:"name"`
	NewStatus    int64            `json:"newStatus"`
	Ov      int64            `json:"ov"`
	Pay     Pay              `json:"pay"`
	Protect int64            `json:"protect"`
	Sa      int64            `json:"sa"`
	Singer  []Album          `json:"singer"`
	Status  int64            `json:"status"`
	Subtitle     string           `json:"subtitle"`
	Tag          int64            `json:"tag"`
	Tid          int64            `json:"tid"`
	TimePublic   string           `json:"time_public"`
	Title        string           `json:"title"`
	TitleHilight string           `json:"title_hilight"`
	Type         int64            `json:"type"`
	URL          string           `json:"url"`
	Version int64            `json:"version"`
	Volume  Volume           `json:"volume"`
}

type Zhida struct {
	Type        int64       `json:"type"`
	ZhidaSinger ZhidaSinger `json:"zhida_singer"`
}

type ZhidaSinger struct {
	AlbumNum          int64         `json:"albumNum"`
	Hotalbum          []interface{} `json:"hotalbum"`
	Hotsong           []interface{} `json:"hotsong"`
	MvNum             int64         `json:"mvNum"`
	SingerID          int64         `json:"singerID"`
	SingerMID         string        `json:"singerMID"`
	SingerName        string        `json:"singerName"`
	SingerPic         string        `json:"singerPic"`
	SingernameHilight string        `json:"singername_hilight"`
	SongNum           int64         `json:"songNum"`
}
