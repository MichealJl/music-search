package qq

import "encoding/json"

type GetPlayURLRsp struct {
	Code    int64  `json:"code"`
	Ts      int64  `json:"ts"`
	StartTs int64  `json:"start_ts"`
	Req0    QQReq0 `json:"req_0"`
}

type QQReq0 struct {
	Code int64            `json:"code"`
	Data QQGetPlayURLData `json:"data"`
}

type QQGetPlayURLData struct {
	Uin          string       `json:"uin"`
	Retcode      int64        `json:"retcode"`
	VerifyType   int64        `json:"verify_type"`
	LoginKey     string       `json:"login_key"`
	Msg          string       `json:"msg"`
	SIP          []string     `json:"sip"`
	Thirdip      []string     `json:"thirdip"`
	Testfile2G   string       `json:"testfile2g"`
	Testfilewifi string       `json:"testfilewifi"`
	Midurlinfo   []Midurlinfo `json:"midurlinfo"`
	Servercheck  string       `json:"servercheck"`
	Expiration   int64        `json:"expiration"`
}

type Midurlinfo struct {
	Songmid           string `json:"songmid"`
	Filename          string `json:"filename"`
	Purl              string `json:"purl"`
	Errtype           string `json:"errtype"`
	P2Pfromtag        int64  `json:"p2pfromtag"`
	Qmdlfromtag       int64  `json:"qmdlfromtag"`
	CommonDownfromtag int64  `json:"common_downfromtag"`
	VipDownfromtag    int64  `json:"vip_downfromtag"`
	Pdl               int64  `json:"pdl"`
	Premain           int64  `json:"premain"`
	Hisdown           int64  `json:"hisdown"`
	Hisbuy            int64  `json:"hisbuy"`
	UIAlert           int64  `json:"uiAlert"`
	Isbuy             int64  `json:"isbuy"`
	Pneedbuy          int64  `json:"pneedbuy"`
	Pneed             int64  `json:"pneed"`
	Isonly            int64  `json:"isonly"`
	Onecan            int64  `json:"onecan"`
	Result            int64  `json:"result"`
	Tips              string `json:"tips"`
	Opi48Kurl         string `json:"opi48kurl"`
	Opi96Kurl         string `json:"opi96kurl"`
	Opi192Kurl        string `json:"opi192kurl"`
	Opiflackurl       string `json:"opiflackurl"`
	Opi128Kurl        string `json:"opi128kurl"`
	Opi192Koggurl     string `json:"opi192koggurl"`
	Wififromtag       string `json:"wififromtag"`
	Flowfromtag       string `json:"flowfromtag"`
	Wifiurl           string `json:"wifiurl"`
	Flowurl           string `json:"flowurl"`
	Vkey              string `json:"vkey"`
	Opi30Surl         string `json:"opi30surl"`
	Ekey              string `json:"ekey"`
	AuthSwitch        int64  `json:"auth_switch"`
	Subcode           int64  `json:"subcode"`
	Opi96Koggurl      string `json:"opi96koggurl"`
}

func (r *GetPlayURLDataParam) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type GetPlayURLDataParam struct {
	Req0 Req0 `json:"req_0"`
}

type Req0 struct {
	Module string `json:"module"`
	Method string `json:"method"`
	Param  Param  `json:"param"`
}

type Param struct {
	GUID      string   `json:"guid"`
	Songmid   []string `json:"songmid"`
	Filename  []string `json:"filename"`
	Songtype  []int64  `json:"songtype"`
	Uin       string   `json:"uin"`
	Loginflag int64    `json:"loginflag"`
	Platform  string   `json:"platform"`
}
