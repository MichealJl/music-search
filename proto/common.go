package proto

const (
	DefaultPage     = 1
	DefaultPageSize = 20
	DefaultCoverUrl = "http://static.emu999.net/pc/images/no-down-ico.png"
)

type SearchParams struct {
	Keyword  string `json:"keyword"`
	Page     uint   `json:"page"`
	PageSize uint   `json:"page_size"`
}

type SearchRet struct {
	List []SearchRsItem `json:"list"`
}

type SearchRsItem struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Artist   string `json:"artist"`
	Album    string `json:"album"`
	PicID    string `json:"pic_id"`
	URLID    string `json:"url_id"`
	LyricID  string `json:"lyric_id"`
	Duration int64  `json:"duration"`
	CoverUrl string `json:"cover_url"`
	Source   string `json:"source"`
}

type GetPlayUrlRet struct {
	Url  string `json:"url"`
	Size int64  `json:"size"`
}
