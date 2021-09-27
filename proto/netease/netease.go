package netease

type SearchNeteaseRsp struct {
	Result Result `json:"result"`
	Code   int64  `json:"code"`
}

type Result struct {
	Songs     []Song `json:"songs"`
	SongCount int64  `json:"songCount"`
}

type Song struct {
	ID          int64       `json:"id"`
	Name        string      `json:"name"`
	Artists     []Artist    `json:"artists"`
	Album       Album       `json:"album"`
	Duration    int64       `json:"duration"`
	CopyrightID int64       `json:"copyrightId"`
	Status      int64       `json:"status"`
	Alias       []string    `json:"alias"`
	Rtype       int64       `json:"rtype"`
	Ftype       int64       `json:"ftype"`
	Mvid        int64       `json:"mvid"`
	Fee         int64       `json:"fee"`
	RURL        interface{} `json:"rUrl"`
	Mark        int64       `json:"mark"`
}

type Album struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	Artist      Artist   `json:"artist"`
	PublishTime int64    `json:"publishTime"`
	Size        int64    `json:"size"`
	CopyrightID int64    `json:"copyrightId"`
	Status      int64    `json:"status"`
	PicID       int64    `json:"picId"`
	Mark        int64    `json:"mark"`
	Alia        []string `json:"alia,omitempty"`
}

type Artist struct {
	ID        int64       `json:"id"`
	Name      string      `json:"name"`
	PicURL    string      `json:"picUrl"`
	Alias     []string    `json:"alias"`
	AlbumSize int64       `json:"albumSize"`
	PicID     int64       `json:"picId"`
	Img1V1URL string      `json:"img1v1Url"`
	Img1V1    int64       `json:"img1v1"`
	Trans     interface{} `json:"trans"`
}

type LyricRsp struct {
	Sgc       bool      `json:"sgc"`
	Sfy       bool      `json:"sfy"`
	Qfy       bool      `json:"qfy"`
	Lrc       Lrc       `json:"lrc"`
	Code      int64     `json:"code"`
}

type Lrc struct {
	Version int64  `json:"version"`
	Lyric   string `json:"lyric"`
}
