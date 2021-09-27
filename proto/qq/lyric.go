package qq

type LyricRsp struct {
	Retcode int64  `json:"retcode"`
	Code    int64  `json:"code"`
	Subcode int64  `json:"subcode"`
	Lyric   string `json:"lyric"`
	Trans   string `json:"trans"`
}
